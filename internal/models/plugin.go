package models

type PluginName string

type PluginResult struct {
	// PluginName record where this result comes from
	PluginName PluginName

	// Data is where plugin store output
	Data any

	// Render string
}

type Plugin interface {
	// Name return the plugin name for calling on manager
	Name() PluginName

	// Process function process [Event] with all original data read-only
	// It return [Result] and error
	//
	// WARN: All plugins should not change any original data on the database
	// Plugin can only alter their own table on database
	Process(e *Event) (PluginResult, error)
}
