package ohauthcat

// AccessPolicy is the value of Profile.AccessPolicy.
type AccessPolicy string

// AccessPolicyGuest marks a sign-in admitted under the organization's guest
// policy: an email domain outside its verified domains.
const AccessPolicyGuest AccessPolicy = "guest"

// IsGuest reports whether the profile signed in under the guest policy.
// EmailDomainVerified stays true for guests, so use this to tell them apart.
func (p Profile) IsGuest() bool {
	return p.AccessPolicy != nil && AccessPolicy(*p.AccessPolicy) == AccessPolicyGuest
}
