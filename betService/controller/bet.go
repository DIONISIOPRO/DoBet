package controller

import (
	"net/http"
	"strconv"

	"github.com/dionisiopro/dobet-bet/domain/bet"
	"github.com/gin-gonic/gin"
)

type BetCreationSucessResponse struct {
	Id string `json:"bet_id"`
}
type BetListResponse struct {
	Bets []bet.BetBase `json:"bets"`
}

type BetResponseError struct {
	Msg string `json:"msg"`
}

type BetService interface {
	CreateBet(bet *bet.BetBase) (string, error)
	BetByUser(user_id string, page, perpage int64) ([]bet.BetBase, error)
	Bets(page, perpage int64) ([]bet.BetBase, error)
}

type BetController struct {
	betService BetService
}

func NewBetController(betService BetService) *BetController {
	return &BetController{
		betService: betService,
	}
}

// GetBets godoc
// @Summary get all bet in the system << only for admin>>
// @Description if you are admin, this route allows you to to get bets in the dobet server
// @Accept  json
// @Produce  json
// @Success 200 {object} BetListResponse	"This doc give a list of bets"
// @Failure 500 {object} BetResponseError    "this is the msg of the error occured"
// @Param    page query     int    false "page"
// @Param    perpage query     int    false "perpage"
// @Router /bets [get]
func (controller *BetController) GetBets() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil {
			perpage = 0
		}
		bets, err := controller.betService.Bets(int64(page), int64(perpage))
		if err != nil {
			c.JSON(http.StatusBadRequest, BetResponseError{Msg: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"bets": BetListResponse{Bets: bets}})
	}
}

// GetBetsById godoc
// @Summary Get the bets by user id
// @Description This route allows you to fetch  by user id
// @Accept  json
// @Produce  json
// @Success 200 {object} BetListResponse	"you will receive the bets in this document"
// @Failure 500 {object} BetResponseError  "this doc return a description of the error occured"
// @Param      page      query    int  false "give the page number"
// @Param       perpage    query    int	  false "give how many elements you want per page"
// @Param       id  path    string  true "user id"
// @Router /bets/:id [get]
func (controller *BetController) GetBetsByUserId() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil {
			page = 0
		}
		perpage, err := strconv.Atoi(c.Query("perpage"))
		if err != nil {
			perpage = 0
		}
		id := c.Param("id")
		bets, err := controller.betService.BetByUser(id, int64(page), int64(perpage))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"bets": bets})
	}
}

// CreateBet godoc
// @Summary Place and Submit your bet
// @Description This route allows your to place your bet
// @Accept  json
// @Produce  json
// @Success 200 {object} BetCreationSucessResponse	"this doc give you succes msg"
// @Failure 500 {object} BetResponseError  "this doc give you a possible has occured"
// @Param       {object}    body    bet.BetBase	  true "Pleasse provide a valid bet document"
// @Router /bet [post]
func (controller *BetController) CreateBet() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bet bet.BetBase
		if err := c.BindJSON(&bet); err != nil {
			msg := "Error while creating the bet, please provide a valid bet"
			c.JSON(http.StatusBadRequest, BetResponseError{Msg: msg})
			c.Abort()
			return
		}
		bet_id, err := controller.betService.CreateBet(&bet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, BetResponseError{Msg: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, BetCreationSucessResponse{Id: bet_id})
		c.Abort()
	}
}
