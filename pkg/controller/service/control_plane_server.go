// Copyright 2020 The Atlas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller_service

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"github.com/binarly-io/atlas/pkg/api/atlas"
	"github.com/binarly-io/atlas/pkg/controller"
)

func GetOutgoingQueue() chan *atlas.Command {
	return controller.GetOutgoing()
}

func GetIncomingQueue() chan *atlas.Command {
	return controller.GetIncoming()
}

type ControlPlaneServer struct {
	atlas.UnimplementedControlPlaneServer
}

func NewControlPlaneServer() *ControlPlaneServer {
	return &ControlPlaneServer{}
}

// CommandsHandler is a user-provided handler for Commands call
var CommandsHandler func(atlas.ControlPlane_CommandsServer, jwt.MapClaims) error

// Commands gRPC call
func (s *ControlPlaneServer) Commands(CommandsServer atlas.ControlPlane_CommandsServer) error {
	log.Info("Commands() - start")
	defer log.Info("Commands() - end")

	if CommandsHandler == nil {
		return fmt.Errorf("no CommandsHandler provided")
	}

	metadata := fetchMetadata(CommandsServer.Context())
	return CommandsHandler(CommandsServer, metadata)

	// controller.CommandsExchangeEndlessLoop(CommandsServer)
	// return nil
}

// DataChunksHandler is a user-provided handler for DataChunks call
var DataChunksHandler func(atlas.ControlPlane_DataChunksServer, jwt.MapClaims) error

// DataChunks gRPC call
func (s *ControlPlaneServer) DataChunks(DataChunksServer atlas.ControlPlane_DataChunksServer) error {
	log.Info("DataChunks() - start")
	defer log.Info("DataChunks() - end")

	if DataChunksHandler == nil {
		return fmt.Errorf("no DataChunksHandler provided")
	}

	metadata := fetchMetadata(DataChunksServer.Context())
	return DataChunksHandler(DataChunksServer, metadata)
}

// FileStatusHandler is a user-provided handler for FileStatus call
var FileStatusHandler func(*atlas.FileFingerprint, jwt.MapClaims) (*atlas.FileStatus1, error)

// FileStatus gRPC call
func (s *ControlPlaneServer) FileStatus(ctx context.Context, req *atlas.FileFingerprint) (*atlas.FileStatus1, error) {
	log.Info("FileStatus() - start")
	defer log.Info("FileStatus() - end")

	if FileStatusHandler == nil {
		return nil, fmt.Errorf("no FileStatusHandler provided")
	}

	metadata := fetchMetadata(ctx)
	return FileStatusHandler(req, metadata)
}

