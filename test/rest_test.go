/*
 *  *******************************************************************************
 *  * Copyright (c) 2019 Edgeworx, Inc.
 *  *
 *  * This program and the accompanying materials are made available under the
 *  * terms of the Eclipse Public License v. 2.0 which is available at
 *  * http://www.eclipse.org/legal/epl-2.0
 *  *
 *  * SPDX-License-Identifier: EPL-2.0
 *  *******************************************************************************
 *
 */

package resttest

import (
	"fmt"
	"net/url"
	"testing"

	"github.com/datasance/iofog-go-sdk/v3/pkg/client"
)

type testState struct {
	email           string
	password        string
	name            string
	surname         string
	url             *url.URL
	agent           string
	uuid            string
	fogType         int64
	appTemplateName string
	appName         string
}

var state = testState{
	email:           "serge@edgeworx.io",
	password:        "wfhoi982bv1sfdjoi",
	name:            "Serge",
	surname:         "Radinovich",
	agent:           "agent-1",
	fogType:         1, // x86
	appTemplateName: "apptemplate1",
	appName:         "app-1",
}

var clt *client.Client

func TestNewAndLogin(t *testing.T) {
	// client.SetVerbosity(true)

	var existingState = testState{
		email:    "user@domain.com",
		password: "g9hr823rhuoi",
		name:     "Foo",
		surname:  "Bar",
	}
	var err error
	state.url, err = url.Parse("http://localhost:51121/api/v3")
	if err != nil {
		t.Error(err)
	}
	opt := client.Options{
		BaseURL: state.url,
	}

	clt, err := client.NewAndLogin(opt, existingState.email, existingState.password)
	if err != nil {
		t.Fatalf(fmt.Sprintf("Failed to create client and login: %s", err.Error()))
	}

	_, err = clt.GetStatus()
	if err != nil {
		t.Fatalf(fmt.Sprintf("Failed to get status: %s", err.Error()))
	}
}

func TestNewAndCreate(t *testing.T) {
	opt := client.Options{
		BaseURL: state.url,
	}
	clt = client.New(opt)

	if err := clt.CreateUser(client.User{
		Email:    state.email,
		Password: state.password,
		Name:     state.name,
		Surname:  state.surname,
	}); err != nil {
		t.Fatalf(fmt.Sprintf("Failed to create user : %s", err.Error()))
	}

	_, err := clt.GetStatus()
	if err != nil {
		t.Fatalf(fmt.Sprintf("Failed to get status: %s", err.Error()))
	}

	if err = clt.Login(client.LoginRequest{
		Email:    state.email,
		Password: state.password,
	}); err != nil {
		t.Fatalf(fmt.Sprintf("Failed to login: %s", err.Error()))
	}
}

func TestCreateAgent(t *testing.T) {
	request := &client.CreateAgentRequest{}
	request.FogType = &state.fogType
	request.Name = state.agent
	host := "localhost"
	request.Host = &host

	response, err := clt.CreateAgent(request)
	if err != nil {
		t.Fatalf(fmt.Sprintf("Failed to create Agent: %s", err.Error()))
	}

	getResponse, err := clt.GetAgentByID(response.UUID)
	if err != nil {
		t.Fatalf((fmt.Sprintf("Failed to get Agent by UUID: %s", err.Error())))
	}

	if getResponse.Name != request.Name {
		t.Fatalf(fmt.Sprintf("Controller returned unexpected Agent name: %s", getResponse.Name))
	}

	nameInfo, err := clt.GetAgentByName(state.agent, false)
	if err != nil {
		t.Fatalf("Failed to get Agent by name: %s", err.Error())
	}
	idInfo, err := clt.GetAgentByID(nameInfo.UUID)
	if err != nil {
		t.Fatalf("Failed to get Agent by UUID: %s", err.Error())
	}
	state.uuid = idInfo.UUID
}

func TestDeleteAgent(t *testing.T) {
	if err := clt.DeleteAgent(state.uuid); err != nil {
		t.Fatalf("Failed to delete Agent: %s", err.Error())
	}
}
