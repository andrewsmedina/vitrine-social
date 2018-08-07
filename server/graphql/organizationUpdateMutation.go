package graphql

import (
	"errors"

	"github.com/Coderockr/vitrine-social/server/model"
	"github.com/gobuffalo/pop/nulls"
	"github.com/graphql-go/graphql"
)

type updateOrgFn func(model.Organization) (model.Organization, error)

var addressInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AddressInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"street":        stringInput,
		"number":        stringInput,
		"complement":    stringInput,
		"neighbordhood": stringInput,
		"city":          stringInput,
		"state":         stringInput,
		"zipcode":       stringInput,
		"withoutComplement": &graphql.InputObjectFieldConfig{
			Type:         graphql.Boolean,
			Description:  "When set to true, will make the address without a complement",
			DefaultValue: false,
		},
	},
})

var organizationUpdateInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OrganizationUpdateInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name":  stringInput,
		"phone": stringInput,
		"about": stringInput,
		"video": stringInput,
		"email": stringInput,
		"address": &graphql.InputObjectFieldConfig{
			Type: addressInput,
		},
	},
})

var organizationUpdatePayload = graphql.NewObject(graphql.ObjectConfig{
	Name: "OrganizationUpdatePayload",
	Fields: graphql.Fields{
		"organization": &graphql.Field{
			Type: organizationType,
		},
	},
})

func newOrganizationUpdateMutation(update updateOrgFn) *graphql.Field {
	return &graphql.Field{
		Name:        "OrganizationUpdateMutation",
		Description: "Updates Organization Profile",
		Args: graphql.FieldConfigArgument{
			"input": &graphql.ArgumentConfig{
				Type: organizationUpdateInput,
			},
		},
		Type: organizationUpdatePayload,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			o, ok := p.Source.(*model.Organization)
			if !ok || o == nil {
				return nil, errors.New("organization not found")
			}

			input := p.Args["input"].(map[string]interface{})

			if name, ok := input["name"].(string); ok {
				o.Name = name
			}

			if about, ok := input["about"].(string); ok {
				o.About = about
			}

			if phone, ok := input["phone"].(string); ok {
				o.Phone = phone
			}

			if video, ok := input["video"].(string); ok {
				o.Video = video
			}

			if email, ok := input["email"].(string); ok {
				o.Email = email
			}

			if address, ok := input["address"].(map[string]interface{}); ok {
				if street, ok := address["street"].(string); ok {
					o.Address.Street = street
				}

				if number, ok := address["number"].(string); ok {
					o.Address.Number = number
				}

				if neighbordhood, ok := address["neighbordhood"].(string); ok {
					o.Address.Neighborhood = neighbordhood
				}

				if city, ok := address["city"].(string); ok {
					o.Address.City = city
				}

				if state, ok := address["state"].(string); ok {
					o.Address.State = state
				}

				if zipcode, ok := address["zipcode"].(string); ok {
					o.Address.Zipcode = zipcode
				}

				withoutComplement := address["withoutComplement"].(bool)
				if withoutComplement {
					o.Address.Complement = nulls.String{Valid: false}
				}

				if complement, ok := address["complement"].(string); ok {
					if withoutComplement {
						return nil, errors.New("parameters withoutComplement and complement can't be used together")
					}
					o.Address.Complement = nulls.NewString(complement)
				}
			}

			var err error
			*o, err = update(*o)
			if err != nil {
				return nil, err
			}

			return map[string]interface{}{"organization": o}, nil
		},
	}
}