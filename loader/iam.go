package loader

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	iamTypes "github.com/aws/aws-sdk-go-v2/service/iam/types"
	log "github.com/sirupsen/logrus"
)

func FetchAllUsers(
	ctx context.Context,
	cfg aws.Config,
) ([]iamTypes.User, error) {
	log.Debug("Fetching all IAM users")
	users := []iamTypes.User{}
	client := iam.NewFromConfig(cfg)
	load := func(nextToken *string) (*string, error) {
		result, err := client.ListUsers(ctx, &iam.ListUsersInput{Marker: nextToken})
		if err != nil {
			return nil, err
		}
		users = append(users, result.Users...)
		return result.Marker, nil
	}
	err := FetchAll("users", load)
	if err != nil {
		return nil, err
	}
	log.Infof("Fetched %d IAM users", len(users))
	return users, nil
}

func FetchAllRoles(
	ctx context.Context,
	cfg aws.Config,
) ([]iamTypes.Role, error) {
	log.Debug("Fetching all IAM roles")
	roles := []iamTypes.Role{}
	client := iam.NewFromConfig(cfg)
	load := func(nextToken *string) (*string, error) {
		result, err := client.ListRoles(ctx, &iam.ListRolesInput{
			Marker: nextToken,
		})
		if err != nil {
			return nil, err
		}
		roles = append(roles, result.Roles...)
		return result.Marker, nil
	}
	err := FetchAll("roles", load)
	if err != nil {
		return nil, err
	}
	log.Infof("Fetched %d IAM roles", len(roles))
	return roles, nil
}

func FetchAllGroups(
	ctx context.Context,
	cfg aws.Config,
) ([]iamTypes.Group, error) {
	log.Debug("Fetching all IAM groups")
	groups := []iamTypes.Group{}
	client := iam.NewFromConfig(cfg)
	load := func(nextToken *string) (*string, error) {
		result, err := client.ListGroups(ctx, &iam.ListGroupsInput{Marker: nextToken})
		if err != nil {
			return nil, err
		}
		groups = append(groups, result.Groups...)
		return result.Marker, nil
	}
	err := FetchAll("groups", load)
	if err != nil {
		return nil, err
	}
	log.Infof("Fetched %d IAM groups", len(groups))
	return groups, nil
}

func FetchAllPolicies(
	ctx context.Context,
	cfg aws.Config,
) ([]iamTypes.Policy, error) {
	log.Debug("Fetching all IAM policies")
	policies := []iamTypes.Policy{}
	client := iam.NewFromConfig(cfg)
	load := func(nextToken *string) (*string, error) {
		result, err := client.ListPolicies(ctx, &iam.ListPoliciesInput{
			Marker: nextToken,
		})
		if err != nil {
			return nil, err
		}
		policies = append(policies, result.Policies...)
		return result.Marker, nil
	}
	err := FetchAll("policies", load)
	if err != nil {
		return nil, err
	}
	log.Infof("Fetched %d IAM policies", len(policies))
	return policies, nil
}

func FetchAllUserGroups(
	ctx context.Context,
	cfg aws.Config,
	user string,
) ([]iamTypes.Group, error) {
	log.Debugf("Fetching all IAM groups for %s", user)
	groups := []iamTypes.Group{}
	client := iam.NewFromConfig(cfg)
	load := func(nextToken *string) (*string, error) {
		result, err := client.ListGroupsForUser(ctx, &iam.ListGroupsForUserInput{
			Marker:   nextToken,
			UserName: &user,
		})
		if err != nil {
			return nil, err
		}
		groups = append(groups, result.Groups...)
		return result.Marker, nil
	}
	err := FetchAll("groups", load)
	if err != nil {
		return nil, err
	}
	log.Debugf("Fetched %d IAM groups for %s", len(groups), user)
	return groups, nil
}

func FetchAllAccessKeys(
	ctx context.Context,
	cfg aws.Config,
	user string,
) ([]iamTypes.AccessKeyMetadata, error) {
	log.Debugf("Fetching all IAM access keys for %s", user)
	accessKeys := []iamTypes.AccessKeyMetadata{}
	client := iam.NewFromConfig(cfg)
	load := func(nextToken *string) (*string, error) {
		result, err := client.ListAccessKeys(ctx, &iam.ListAccessKeysInput{
			Marker:   nextToken,
			UserName: &user,
		})
		if err != nil {
			return nil, err
		}
		accessKeys = append(accessKeys, result.AccessKeyMetadata...)
		return result.Marker, nil
	}
	err := FetchAll("accessKeys", load)
	if err != nil {
		return nil, err
	}
	log.Debugf("Fetched %d IAM access keys for user %s", len(accessKeys), user)
	return accessKeys, nil
}
