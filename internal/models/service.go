package models

// Service hold plugins and database
type Service struct {
	allPlugins map[PluginName]Plugin

	// enabledPlugins is a sequence slice makes sure plugins work properly
	enabledPlugins []Plugin
}

// Register register new plugins for extension
// See the definition of [Plugin] for implement
// And call [Service.ProcessWith] Function to process event with specific plugin
func (s *Service) Register(p Plugin) {
	s.allPlugins[p.Name()] = p
}

// Enable enable a plugin, when [Service.Process] is called, event will be processed by all enabled plugins
func (s *Service) Enable(pn PluginName) error {
	plugin, ok := s.allPlugins[pn]
	if !ok {
		return ErrNotFound
	}

	s.enabledPlugins = append(s.enabledPlugins, plugin)

	return nil
}

// Process process this event with all enabled plugins
// if no error reported, return nil
// else return []error
func (s *Service) Process(e *Event) ([]*Facet, []error) {
	allRes := make([]*Facet, 0, len(s.enabledPlugins))
	allErr := make([]error, 0, 2)

	// traverse all plugins then call their Process functions
	for _, plugin := range s.enabledPlugins {

		r, err := plugin.Process(e)
		if err != nil {
			// if occur error, skip this result
			allErr = append(allErr, err)
			continue
		}
		allRes = append(allRes, r)
	}
	if len(allErr) != 0 {
		return allRes, allErr
	}

	return allRes, nil
}

// ProcessWith process this event with single Plugin by [PluginName]
// It will call [Plugin.Process] function then return error
func (s *Service) ProcessWith(pn PluginName, e *Event) (*Facet, error) {
	return s.allPlugins[pn].Process(e)
}
