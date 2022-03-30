package service

type LogoutStateManager struct {
	stateStore map[string]bool
}
 var logoutManager = new(LogoutStateManager)
func NewLogoutMangger() *LogoutStateManager{
	store := map[string]bool{}
	logoutManager.stateStore = store
	return logoutManager
}

func (manager LogoutStateManager) Logout(id string) {
	manager.stateStore["id"] = false
}

func (manager LogoutStateManager) LogIn(id string) {
	manager.stateStore["id"] = true
}

func (manager LogoutStateManager) IsLogIn(id string) bool {
	login, ok := manager.stateStore["id"]
	return ok && login
}
