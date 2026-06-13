package actor

// defaultSystem 是进程级全局 ActorSystem 单例。
var defaultSystem = NewActorSystem()

// Default 返回全局 ActorSystem 单例。
func Default() *ActorSystem { return defaultSystem }

// Spawn 在全局系统中注册并启动一个 Actor。
func Spawn(name string, a Actor, opts ...SpawnOption) PID {
	return defaultSystem.Spawn(name, a, opts...)
}

// Send 通过全局系统向目标 Actor 发送消息。
func Send(pid PID, msg interface{}) bool { return defaultSystem.Send(pid, msg) }

// Lookup 在全局系统中按名称查找 Actor。
func Lookup(name string) (PID, bool) { return defaultSystem.Lookup(name) }

// MustLookup 在全局系统中按名称查找 Actor，未找到则 panic。
func MustLookup(name string) PID {
	return defaultSystem.MustLookup(name)
}

// Request 通过全局系统向目标 Actor 发送请求并返回 Future。
func Request(pid PID, msg interface{}) *Future { return defaultSystem.Request(pid, msg) }
