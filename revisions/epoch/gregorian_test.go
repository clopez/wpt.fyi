// +build small

// Copyright 2018 The WPT Dashboard Project. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package epoch_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/web-platform-tests/wpt.fyi/revisions/epoch"
)

var monthly epoch.Epoch
var weekly epoch.Epoch
var daily epoch.Epoch
var hourly epoch.Epoch
var oneHourEastOfUTC *time.Location
var oneHourWestOfUTC *time.Location

func init() {
	monthly = epoch.Monthly{}
	weekly = epoch.Weekly{}
	daily = epoch.Daily{}
	hourly = epoch.Hourly{}
	oneHourEastOfUTC = time.FixedZone("UTC+1", 60*60)
	oneHourWestOfUTC = time.FixedZone("UTC-1", -60*60)
}

//
// Monthly
//

func TestIsMonthly_Close(t *testing.T) {
	justPrior := time.Date(2018, 3, 31, 23, 59, 59, 999999999, time.UTC)
	justAfter := time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)
	assert.True(t, monthly.IsEpochal(justPrior, justAfter))
	assert.True(t, monthly.IsEpochal(justAfter, justPrior))
}

func TestIsMonthly_Far(t *testing.T) {
	lastYear := time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC)
	thisYear := time.Date(2018, 5, 1, 0, 0, 0, 0, time.UTC)
	assert.True(t, monthly.IsEpochal(lastYear, thisYear))
	assert.True(t, monthly.IsEpochal(thisYear, lastYear))
}

func TestIsMonthly_TZ(t *testing.T) {
	justPrior := time.Date(2018, 3, 31, 23, 59, 59, 999999999, time.UTC)
	justAfter := time.Date(2018, 3, 31, 23, 59, 59, 999999999, oneHourWestOfUTC)
	assert.True(t, monthly.IsEpochal(justPrior, justAfter))
	assert.True(t, monthly.IsEpochal(justAfter, justPrior))
}

func TestIsNotMonthly_Close(t *testing.T) {
	lastInstant := time.Date(2018, 5, 1, 0, 0, 0, 0, time.UTC)
	nextInstant := time.Date(2018, 5, 1, 0, 0, 0, 1, time.UTC)
	assert.False(t, monthly.IsEpochal(lastInstant, nextInstant))
	assert.False(t, monthly.IsEpochal(nextInstant, lastInstant))
}

func TestIsNotMonthly_Far(t *testing.T) {
	monthStart := time.Date(2018, 3, 1, 0, 0, 0, 0, time.UTC)
	monthEnd := time.Date(2018, 3, 31, 23, 59, 59, 999999999, time.UTC)
	assert.False(t, monthly.IsEpochal(monthStart, monthEnd))
	assert.False(t, monthly.IsEpochal(monthEnd, monthStart))
}

func TestIsNotMonthly_TZ(t *testing.T) {
	monthStart := time.Date(2018, 3, 1, 0, 0, 0, 0, time.UTC)
	monthEnd := time.Date(2018, 4, 1, 0, 0, 0, 0, oneHourEastOfUTC)
	assert.False(t, monthly.IsEpochal(monthStart, monthEnd))
	assert.False(t, monthly.IsEpochal(monthEnd, monthStart))
}

//
// Weekly
//

func TestIsWeekly_Close(t *testing.T) {
	justPrior := time.Date(2018, 4, 1, 23, 59, 59, 999999999, time.UTC)
	justAfter := time.Date(2018, 4, 2, 0, 0, 0, 0, time.UTC)
	assert.True(t, weekly.IsEpochal(justPrior, justAfter))
	assert.True(t, weekly.IsEpochal(justAfter, justPrior))
}

func TestIsWeekly_Far(t *testing.T) {
	lastYear := time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC)
	thisYear := time.Date(2018, 5, 1, 0, 0, 0, 0, time.UTC)
	assert.True(t, weekly.IsEpochal(lastYear, thisYear))
	assert.True(t, weekly.IsEpochal(thisYear, lastYear))
}

func TestIsWeekly_TZ(t *testing.T) {
	justPrior := time.Date(2018, 4, 1, 23, 59, 59, 999999999, time.UTC)
	justAfter := time.Date(2018, 4, 1, 23, 59, 59, 999999999, oneHourWestOfUTC)
	assert.True(t, weekly.IsEpochal(justPrior, justAfter))
	assert.True(t, weekly.IsEpochal(justAfter, justPrior))
}

func TestIsNotWeekly_Close(t *testing.T) {
	lastInstant := time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)
	nextInstant := time.Date(2018, 4, 1, 0, 0, 0, 1, time.UTC)
	assert.False(t, weekly.IsEpochal(lastInstant, nextInstant))
	assert.False(t, weekly.IsEpochal(nextInstant, lastInstant))
}

func TestIsNotWeekly_Far(t *testing.T) {
	weekStart := time.Date(2018, 4, 2, 0, 0, 0, 0, time.UTC)
	weekEnd := time.Date(2018, 4, 8, 23, 59, 59, 999999999, time.UTC)
	assert.False(t, weekly.IsEpochal(weekStart, weekEnd))
	assert.False(t, weekly.IsEpochal(weekEnd, weekStart))
}

func TestIsNotWeekly_TZ(t *testing.T) {
	weekStart := time.Date(2018, 4, 2, 0, 0, 0, 0, time.UTC)
	weekEnd := time.Date(2018, 4, 9, 0, 0, 0, 0, oneHourEastOfUTC)
	assert.False(t, weekly.IsEpochal(weekStart, weekEnd))
	assert.False(t, weekly.IsEpochal(weekEnd, weekStart))
}

//
// Daily
//

func TestIsDaily_Close(t *testing.T) {
	justPrior := time.Date(2018, 4, 1, 23, 59, 59, 999999999, time.UTC)
	justAfter := time.Date(2018, 4, 2, 0, 0, 0, 0, time.UTC)
	assert.True(t, daily.IsEpochal(justPrior, justAfter))
	assert.True(t, daily.IsEpochal(justAfter, justPrior))
}

func TestIsDaily_Far(t *testing.T) {
	lastYear := time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC)
	thisYear := time.Date(2018, 5, 1, 0, 0, 0, 0, time.UTC)
	assert.True(t, daily.IsEpochal(lastYear, thisYear))
	assert.True(t, daily.IsEpochal(thisYear, lastYear))
}

func TestIsDaily_TZ(t *testing.T) {
	justPrior := time.Date(2018, 4, 1, 23, 59, 59, 999999999, time.UTC)
	justAfter := time.Date(2018, 4, 1, 23, 59, 59, 999999999, oneHourWestOfUTC)
	assert.True(t, daily.IsEpochal(justPrior, justAfter))
	assert.True(t, daily.IsEpochal(justAfter, justPrior))
}

func TestIsNotDaily_Close(t *testing.T) {
	lastInstant := time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)
	nextInstant := time.Date(2018, 4, 1, 0, 0, 0, 1, time.UTC)
	assert.False(t, daily.IsEpochal(lastInstant, nextInstant))
	assert.False(t, daily.IsEpochal(nextInstant, lastInstant))
}

func TestIsNotDaily_Far(t *testing.T) {
	dayStart := time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)
	dayEnd := time.Date(2018, 4, 1, 23, 59, 59, 999999999, time.UTC)
	assert.False(t, daily.IsEpochal(dayStart, dayEnd))
	assert.False(t, daily.IsEpochal(dayEnd, dayStart))
}

func TestIsNotDaily_TZ(t *testing.T) {
	dayStart := time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)
	dayEnd := time.Date(2018, 4, 2, 0, 0, 0, 0, oneHourEastOfUTC)
	assert.False(t, daily.IsEpochal(dayStart, dayEnd))
	assert.False(t, daily.IsEpochal(dayEnd, dayStart))
}

//
// Hourly
//

func TestIsHourly_Close(t *testing.T) {
	justPrior := time.Date(2018, 4, 1, 0, 59, 59, 999999999, time.UTC)
	justAfter := time.Date(2018, 4, 2, 1, 0, 0, 0, time.UTC)
	assert.True(t, hourly.IsEpochal(justPrior, justAfter))
	assert.True(t, hourly.IsEpochal(justAfter, justPrior))
}

func TestIsHourly_Far(t *testing.T) {
	lastYear := time.Date(2017, 5, 1, 0, 0, 0, 0, time.UTC)
	thisYear := time.Date(2018, 5, 1, 0, 0, 0, 0, time.UTC)
	assert.True(t, hourly.IsEpochal(lastYear, thisYear))
	assert.True(t, hourly.IsEpochal(thisYear, lastYear))
}

func TestIsHourly_TZ(t *testing.T) {
	justPrior := time.Date(2018, 4, 1, 0, 59, 59, 999999999, time.UTC)
	justAfter := time.Date(2018, 4, 1, 0, 59, 59, 999999999, oneHourWestOfUTC)
	assert.True(t, hourly.IsEpochal(justPrior, justAfter))
	assert.True(t, hourly.IsEpochal(justAfter, justPrior))
}

func TestIsNotHourly_Close(t *testing.T) {
	lastInstant := time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)
	nextInstant := time.Date(2018, 4, 1, 0, 0, 0, 1, time.UTC)
	assert.False(t, hourly.IsEpochal(lastInstant, nextInstant))
	assert.False(t, hourly.IsEpochal(nextInstant, lastInstant))
}

func TestIsNotHourly_Far(t *testing.T) {
	hourStart := time.Date(2018, 4, 1, 1, 0, 0, 0, time.UTC)
	hourEnd := time.Date(2018, 4, 1, 1, 59, 59, 999999999, time.UTC)
	assert.False(t, hourly.IsEpochal(hourStart, hourEnd))
	assert.False(t, hourly.IsEpochal(hourEnd, hourStart))
}

func TestIsNotHourly_TZ(t *testing.T) {
	hourStart := time.Date(2018, 4, 1, 1, 0, 0, 0, time.UTC)
	hourEnd := time.Date(2018, 4, 1, 2, 0, 0, 0, oneHourEastOfUTC)
	assert.False(t, hourly.IsEpochal(hourStart, hourEnd))
	assert.False(t, hourly.IsEpochal(hourEnd, hourStart))
}
