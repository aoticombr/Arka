package main

import "github.com/gin-gonic/gin"

// ExecutorBuilder is a composable builder. Every injected property is used by the NewCmdExecutor method.
type ExecutorBuilder struct {
	// LoggerFactory               LoggerFactory
	// PluginLoader                PluginLoader
	// SubscriberFactoriesRegister SubscriberFactoriesRegister
	// TokenRejecterFactory        TokenRejecterFactory
	// MetricsAndTracesRegister    MetricsAndTracesRegister
	// EngineFactory               EngineFactory
	// ProxyFactory                ProxyFactory
	// BackendFactory              BackendFactory
	// HandlerFactory              HandlerFactory
	// RunServerFactory            RunServerFactory
	// AgentStarterFactory         AgentStarter

	Middlewares []gin.HandlerFunc
}
