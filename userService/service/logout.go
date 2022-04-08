package service

type LogInStateManager struct {
	stateStore map[string]bool
}
 var logoutManager = new(LogInStateManager)
 
func NewLogInStateManager() *LogInStateManager{
	store := map[string]bool{}
	logoutManager.stateStore = store
	return logoutManager
}

func (manager *LogInStateManager) Logout(id string) {
	manager.stateStore["id"] = false
}

func (manager *LogInStateManager) LogIn(id string) {
	manager.stateStore["id"] = true
}

func (manager *LogInStateManager) IsLogIn(id string) bool {
	login, ok := manager.stateStore["id"]
	return ok && login
}
