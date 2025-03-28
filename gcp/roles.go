type BindingPresence struct {
	NonProdUsers  bool
	NonProdSystem bool
	ProdUsers     bool
	ProdSystem    bool
}

type Governance struct {
	ProjectRoles map[string]RolePolicy
}

type RolePolicy struct {
	Users  EnvPolicyUsers
	System EnvPolicySystem
}

type EnvPolicyUsers struct {
	Non  string // "ambient" or "elevated"
	Prod string // "ambient" or "elevated"
}

type EnvPolicySystem struct {
	Non  bool
	Prod bool
}

// Extract roles with attempted bindings categorized by env and type
func getRoles(app App) map[string]BindingPresence {
	result := make(map[string]BindingPresence)

	for _, rb := range app.Spec.Env.NonProd.RoleBindings {
		bp := result[rb.Role]
		if len(rb.Members["groups"]) > 0 {
			bp.NonProdUsers = true
		}
		if len(rb.Members["SAs"]) > 0 || len(rb.Members["WiF"]) > 0 {
			bp.NonProdSystem = true
		}
		result[rb.Role] = bp
	}

	for _, rb := range app.Spec.Env.Prod.RoleBindings {
		bp := result[rb.Role]
		if len(rb.Members["groups"]) > 0 {
			bp.ProdUsers = true
		}
		if len(rb.Members["SAs"]) > 0 || len(rb.Members["WiF"]) > 0 {
			bp.ProdSystem = true
		}
		result[rb.Role] = bp
	}

	return result
}

// Validate roles against governance rules
func validateRoles(roleMap map[string]BindingPresence, gov Governance) []error {
	var errors []error

	for role, presence := range roleMap {
		policy, exists := gov.ProjectRoles[role]
		if !exists {
			errors = append(errors, fmt.Errorf("role '%s' must be scoped for project level", role))
			continue
		}

		// Users
		if presence.NonProdUsers && policy.Users.Non == "" {
			errors = append(errors, fmt.Errorf("role '%s' is not permitted for non-prod users", role))
		}
		if presence.ProdUsers && policy.Users.Prod == "" {
			errors = append(errors, fmt.Errorf("role '%s' is not permitted for prod users", role))
		}

		// System
		if presence.NonProdSystem && !policy.System.Non {
			errors = append(errors, fmt.Errorf("role '%s' is not permitted for non-prod system accounts", role))
		}
		if presence.ProdSystem && !policy.System.Prod {
			errors = append(errors, fmt.Errorf("role '%s' is not permitted for prod system accounts", role))
		}
	}

	return errors
}
