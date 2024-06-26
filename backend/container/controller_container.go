package container

import controllers "github.com/2110366-2566-2/Mai-Roi-Ra/backend/controllers"

// ControllerProvider Inject Controller
func (c *Container) ControllerProvider() {
	if err := c.Container.Provide(controllers.NewUserController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewTestController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewEventController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewLocationController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewParticipateController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewAnnouncementController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewProblemController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewTransactionController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewRefundController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewPostController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controllers.NewResponseController); err != nil {
		c.Error = err
	}
	
	if err := c.Container.Provide(controllers.NewController); err != nil {
		c.Error = err
	}
	
}
