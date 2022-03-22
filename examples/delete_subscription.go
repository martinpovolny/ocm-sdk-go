/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

// This example shows how to delete a subscription.

import (
	"context"
	"os"

	sdk "github.com/openshift-online/ocm-sdk-go/v2"
)

func deleteSubscription(ctx context.Context, args []string) error {
	// Create the connection, and remember to close it:
	token := os.Getenv("OCM_TOKEN")
	connection, err := sdk.NewConnection().
		Logger(logger).
		Tokens(token).
		BuildContext(ctx)
	if err != nil {
		return err
	}
	defer connection.Close()

	// Get the client for the resource that manages the collection of subscriptions:
	collection := connection.AccountsMgmt().V1().Subscriptions()

	// Get the client for the resource that manages the subscription that we want to delete.
	// Note that this will not send any request to the server yet, so it will succeed even if
	// the subscription doesn't exist.
	resource := collection.Subscription("1BDFg66jv2kDfBh6bBog3IsZWVH")

	// Send the request to delete the subscription:
	_, err = resource.Delete().Send(ctx)
	if err != nil {
		return err
	}

	return nil
}
