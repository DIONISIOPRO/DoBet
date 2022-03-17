package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gitthub.com/dionisiopro/dobet/api"
	"gitthub.com/dionisiopro/dobet/config"
	"gitthub.com/dionisiopro/dobet/controller"
	"gitthub.com/dionisiopro/dobet/database"
	"gitthub.com/dionisiopro/dobet/repository"
	"gitthub.com/dionisiopro/dobet/routes"
	"gitthub.com/dionisiopro/dobet/service"
)

func main() {
	Config := config.BaseConfig{}
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fileDecoder := json.NewDecoder(file)
	err = fileDecoder.Decode(&Config)
	if err != nil {
		panic(err)
	}
	fmt.Print("file redy")
	fmt.Print(Config.App.Host)
	const (
		USERCOLLECTION   = "users"
		ODDCOLLECTION    = "odds"
		LEAGUECOLLECTION = "leagues"
		MATCHCOLLECTION  = "matches"
		TEAMCOLLECTION   = "teams"
		BETCOLLECTION    = "bets"
	)
	var client = &http.Client{}
	var footballApi = api.NewFootBallApi(client, Config.Api.BaseUrl, Config.Api.Token, Config.Api.Host)

	var userCollection = database.OpenCollection(USERCOLLECTION)
	var oddCollection = database.OpenCollection(ODDCOLLECTION)
	var leagueCollection = database.OpenCollection(LEAGUECOLLECTION)
	var matchCollection = database.OpenCollection(MATCHCOLLECTION)
	var teamCollection = database.OpenCollection(TEAMCOLLECTION)
	var betCollection = database.OpenCollection(BETCOLLECTION)
	fmt.Print("collections created")

	var userRepository = repository.NewUserRepository(userCollection)
	var oddRepository = repository.NewOddRepository(oddCollection)
	var leagueRepository = repository.NewLeagueRepository(leagueCollection)
	var matchRepository = repository.NewMatchReposiotry(matchCollection)
	var teamRepository = repository.NewTeamRepository(teamCollection)
	var paymentRepository = repository.NewPaymentReposiotry(userCollection)
	var betRepository = repository.NewBetRepository(paymentRepository, betCollection)
	var authRepository = repository.NewAuthRepository(userCollection)

	fmt.Print("repositories seted")
	var userService = service.NewupUserService(userRepository)
	var leagueService = service.NewLeagueService(leagueRepository, footballApi)
	var oddService = service.NewOddServivce(oddRepository, footballApi, leagueService)
	var betService = service.NewBetService(betRepository)
	var matchService = service.NewMatchService(matchRepository, betService, footballApi, oddService)
	var teamService = service.NewTeamService(teamRepository, footballApi)
	var authService = service.NewAuthService(authRepository)
	var paymentService = service.NewPaymentService(paymentRepository)
	fmt.Print("services created")
	var userController = controller.NewUserController(userService)
	var leagueController = controller.NewLeagueController(leagueService)
	var authController = controller.NewAuthController(authService)
	var betController = controller.NewBetController(betService)
	var matchController = controller.NewMatchController(matchService)
	var paymentController = controller.NewPaymnetController(paymentService)
	var teamController = controller.NewTeamRepository(teamService)
	fmt.Print("controller created")
	var userRouter = routes.NewUserRouter(*userController)
	var leagueRouter = routes.NewLeagueRouter(*leagueController)
	var authRouter = routes.NewAuthRouter(*authController)
	var betRouter = routes.NewBetRouter(*betController)
	var matchRouter = routes.NewMatchRouter(*matchController)
	var paymentRouter = routes.NewPaymentRouter(*paymentController)
	var teamRouter = routes.NewTeamRouter(*teamController)
	fmt.Print("routes created")
	app := gin.New()
	app = userRouter.SetupUserRouter(app)
	app = leagueRouter.SetupLeagueRouter(app)
	app = authRouter.SetupAuthRoutes(app)
	app = betRouter.SetupBetRoutes(app)
	app = matchRouter.SetupMatchRouter(app)
	app = paymentRouter.SetupPaymentRouter(app)
	app = teamRouter.SetupTeamRouter(app)
	fmt.Print("app created")
	go leagueService.LunchUpdateLeaguesLoop()
	go teamService.LunchUpdateTeamssLoop()
	go matchService.MatchWatch()
	go matchService.LunchUpdateMatchesLoop()
	go oddService.LunchUpdateOddsLoop()
	fmt.Print("all goroutines lunched")
	app.Run(fmt.Sprintf("%v:%v", Config.App.Host, Config.App.Port))
}
