package controltower

import (
	"airplane/domain"
	"sync"
)

type controlTower struct {
	isAirStripEmpty bool
	isLandingEnable bool
}

var mw sync.Mutex

func (c *controlTower) SetStatusLanding(b bool) {
	mw.Lock()
	c.isLandingEnable = b
	mw.Unlock()
}

func (c *controlTower) SetStatusAirStrip(b bool) {
	mw.Lock()
	c.isAirStripEmpty = b
	mw.Unlock()
}

func (c *controlTower) IsAirStripEmpty() bool {
	return c.isAirStripEmpty
}

func (c *controlTower) IsLandingEnable() bool {
	return c.isLandingEnable
}

func NewControlTower() domain.ControlTower {
	return &controlTower{}
}
