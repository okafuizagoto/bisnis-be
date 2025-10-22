package agent

import (
	agentEntity "bisnis-be/internal/entity/agent"
	"bisnis-be/pkg/errors"
	"context"
	"fmt"
)

func (d Data) CheckAgent(ctx context.Context, agentUser agentEntity.LoginAgent) (agentEntity.Agent, string, error) {
	var (
		result                  string
		agentIDValidation       agentEntity.Agent
		agentPasswordValidation agentEntity.Agent
		err                     error
	)
	rows, err := (*d.stmt)[getAgentID].QueryxContext(ctx, agentUser.AgentID)
	if err != nil {
		return agentPasswordValidation, result, errors.Wrap(err, "[DATA][CheckAgent][getAgentID]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&agentIDValidation); err != nil {
			return agentPasswordValidation, result, errors.Wrap(err, "[DATA][CheckAgent][getAgentID][Scan]")
		}
	}
	if agentIDValidation != (agentEntity.Agent{}) {
		rowss, err := (*d.stmt)[validationAgentPassword].QueryxContext(ctx, agentUser.AgentID, agentUser.AgentPassword)
		if err != nil {
			return agentPasswordValidation, result, errors.Wrap(err, "[DATA][CheckAgent][validationAgentPassword]")
		}

		defer rowss.Close()

		for rowss.Next() {
			if err = rowss.StructScan(&agentPasswordValidation); err != nil {
				return agentPasswordValidation, result, errors.Wrap(err, "[DATA][CheckAgent][validationAgentPassword][Scan]")
			}
		}
		if agentPasswordValidation == (agentEntity.Agent{}) {
			result = "Incorrect password"
		}
		if agentPasswordValidation != (agentEntity.Agent{}) {
			result = "Success"
		}
	}
	if agentIDValidation == (agentEntity.Agent{}) {
		result = "Agent does not exist"
	}
	return agentPasswordValidation, result, err
}

func (d Data) LoginAgent(ctx context.Context, agentUser agentEntity.LoginAgent) error {
	var (
		err error
	)
	return err
}

func (d Data) AddAgent(ctx context.Context, agentLogin agentEntity.AgentRequest) (string, error) {
	var (
		result            string
		agentIDValidation agentEntity.Agent
		err               error
	)
	fmt.Printf("tsetAgent %+v", agentLogin)
	rows, err := (*d.stmt)[getAgentIDName].QueryxContext(ctx, agentLogin.AgentID, "%"+agentLogin.AgentName+"%")
	if err != nil {
		return result, errors.Wrap(err, "[DATA][AddAgent][getAgentIDName]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&agentIDValidation); err != nil {
			return result, errors.Wrap(err, "[DATA][AddAgent][getAgentIDName][Scan]")
		}
	}

	if agentIDValidation == (agentEntity.Agent{}) {
		_, err = (*d.stmt)[insertAgent].ExecContext(ctx,
			agentLogin.AgentID,
			agentLogin.AgentName,
			agentLogin.AgentPassword,
			agentLogin.Active,
		)
		if err != nil {
			result = "Failed"
			return result, errors.Wrap(err, "[DATA][AddAgent][ExecContext]")
		}
	}
	if agentIDValidation != (agentEntity.Agent{}) {
		return "Agent already registered", errors.Wrap(err, "[DATA][AddAgent]")
	}

	result = "Success"

	return result, err

}

func (d Data) DeleteAgent(ctx context.Context, agentLogin agentEntity.LoginAgent) (agentEntity.Agent, string, error) {
	var (
		result            string
		agentIDValidation agentEntity.Agent
		err               error
	)
	rows, err := (*d.stmt)[getAgentID].QueryxContext(ctx, agentLogin.AgentID)
	if err != nil {
		return agentIDValidation, result, errors.Wrap(err, "[DATA][CheckAgent][getAgentID]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&agentIDValidation); err != nil {
			return agentIDValidation, result, errors.Wrap(err, "[DATA][AddAgent][getAgentIDName][Scan]")
		}
	}

	if agentIDValidation == (agentEntity.Agent{}) {
		return agentIDValidation, "Agent does not exist", errors.Wrap(err, "[DATA][AddAgent]")
	}
	if agentIDValidation != (agentEntity.Agent{}) {
		_, err = (*d.stmt)[deleteAgent].ExecContext(ctx,
			agentLogin.AgentID,
		)
		if err != nil {
			result = "Failed"
			return agentIDValidation, result, errors.Wrap(err, "[DATA][DeleteAgent][ExecContext]")
		}
	}

	result = "Success"

	return agentIDValidation, result, err

}

func (d Data) UpdateAgent(ctx context.Context, agentLogin agentEntity.AgentRequest) (agentEntity.Agent, string, error) {
	var (
		result                string
		agentIDValidation     agentEntity.Agent
		agentIDValidationName agentEntity.Agent
		err                   error
	)
	rows, err := (*d.stmt)[getAgentIDName].QueryxContext(ctx, agentLogin.AgentID, "%"+agentLogin.AgentName+"%")
	if err != nil {
		return agentIDValidation, result, errors.Wrap(err, "[DATA][AddAgent][getAgentIDName]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&agentIDValidation); err != nil {
			return agentIDValidation, result, errors.Wrap(err, "[DATA][AddAgent][getAgentIDName][Scan]")
		}
	}

	rowsz, err := (*d.stmt)[getAgentName].QueryxContext(ctx, "%"+agentLogin.AgentName)
	if err != nil {
		return agentIDValidation, result, errors.Wrap(err, "[DATA][AddAgent][getAgentIDName]")
	}

	defer rowsz.Close()

	for rowsz.Next() {
		if err = rowsz.StructScan(&agentIDValidationName); err != nil {
			return agentIDValidation, result, errors.Wrap(err, "[DATA][AddAgent][getAgentIDName][Scan]")
		}
	}

	if agentIDValidationName != (agentEntity.Agent{}) {
		return agentIDValidation, "Agent Name already registered", errors.Wrap(err, "[DATA][AddAgent]")
	}

	if agentIDValidation == (agentEntity.Agent{}) {
		return agentIDValidation, "Agent does not exist", errors.Wrap(err, "[DATA][AddAgent]")
	}
	if agentIDValidation != (agentEntity.Agent{}) {
		_, err = (*d.stmt)[updateAgent].ExecContext(ctx,
			agentLogin.AgentName,
			agentLogin.AgentID,
		)
		if err != nil {
			result = "Failed"
			return agentIDValidation, result, errors.Wrap(err, "[DATA][DeleteAgent][ExecContext]")
		}
	}

	result = "Success"

	return agentIDValidation, result, err

}
