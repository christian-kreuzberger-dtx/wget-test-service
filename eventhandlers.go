package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	"github.com/keptn-sandbox/sdk-go/pkg/keptn"
)

/**
* Here are all the handler functions for the individual event
  See https://github.com/keptn/spec/blob/0.1.3/cloudevents.md for details on the payload

  -> "sh.keptn.event.configuration.change"
  -> "sh.keptn.events.deployment-finished"
  -> "sh.keptn.events.tests-finished"
  -> "sh.keptn.event.start-evaluation"
  -> "sh.keptn.events.evaluation-done"
  -> "sh.keptn.event.problem.open"
  -> "sh.keptn.events.problem"
*/

//
// Handles ConfigurationChangeEventType = "sh.keptn.event.configuration.change"
// TODO: add in your handler code
//
func HandleConfigurationChangeEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.ConfigurationChangeEventData) error {
	log.Printf("Handling Configuration Changed Event: %s", incomingEvent.Context.GetID())

	return nil
}

//
// Handles DeploymentFinishedEventType = "sh.keptn.events.deployment-finished"
// TODO: add in your handler code
//
func HandleDeploymentFinishedEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.DeploymentFinishedEventData) error {
	log.Printf("Handling Deployment Finished Event: %s", incomingEvent.Context.GetID())

	testsStarted := time.Now()

	// ToDo: Start tests
	// return myKeptn.SendTestsFinishedEvent(&incomingEvent, data.TestStrategy, data.DeploymentStrategy, "TODO", "PASS", data.Labels, "wget-test-service")

	if data.DeploymentURIPublic == "" {
		// no deployment URI available - skip tests
		return myKeptn.SendTestsFinishedEvent(&incomingEvent, "", "", testsStarted, "warning", nil, "wget-test-service")
	}

	log.Printf("Executing wget %s", data.DeploymentURIPublic)
	cmd := exec.Command("wget", data.DeploymentURIPublic)
	stdout, err := cmd.Output()

	fmt.Print(string(stdout))

	if err != nil {
		log.Println(err.Error())

		return myKeptn.SendTestsFinishedEvent(&incomingEvent, "", "", testsStarted, "error", nil, "wget-test-service")
	}

	return myKeptn.SendTestsFinishedEvent(&incomingEvent, "", "", testsStarted, "success", nil, "wget-test-service")
}

//
// Handles TestsFinishedEventType = "sh.keptn.events.tests-finished"
// TODO: add in your handler code
//
func HandleTestsFinishedEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.TestsFinishedEventData) error {
	log.Printf("Handling Tests Finished Event: %s", incomingEvent.Context.GetID())

	return nil
}

//
// Handles EvaluationDoneEventType = "sh.keptn.events.evaluation-done"
// TODO: add in your handler code
//
func HandleStartEvaluationEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.StartEvaluationEventData) error {
	log.Printf("Handling Start Evaluation Event: %s", incomingEvent.Context.GetID())

	return nil
}

//
// Handles DeploymentFinishedEventType = "sh.keptn.events.deployment-finished"
// TODO: add in your handler code
//
func HandleEvaluationDoneEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.EvaluationDoneEventData) error {
	log.Printf("Handling Evaluation Done Event: %s", incomingEvent.Context.GetID())

	return nil
}

//
// Handles ProblemOpenEventType = "sh.keptn.event.problem.open"
// Handles ProblemEventType = "sh.keptn.events.problem"
// TODO: add in your handler code
//
func HandleProblemEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.ProblemEventData) error {
	log.Printf("Handling Problem Event: %s", incomingEvent.Context.GetID())

	return nil
}
