package backend

import (
	"github.com/digitalrebar/digitalrebar/go/common/store"
)

// Profile represents a set of key/values to use in
// template expansion.
//
// There are two types of special profile lists
// global and per machine.
//
// global is named 'global'
// per-machine are marked machine-only and named after the machine.
//
// These can be assigned to a machine's profile list.
// swagger:model
type Profile struct {
	// The name of the profile.  THis must be unique across all
	// profiles. By convention for machine specific,
	// it is the FQDN of the machine, these are created at machine
	// construction and deletion time.
	//
	// required: true
	// swagger:strfmt hostname
	Name string
	// A description of this profile.  This can contain any reference
	// information for humans you want associated with the profile.
	Description string
	// Any additional parameters that may be needed to expand templates
	// for BootEnv, as documented by that boot environment's
	// RequiredParams and OptionalParams.
	Params map[string]interface{}
	//
	// Is this a machine-only, profile
	MachineOnly bool

	p *DataTracker
}

func (p *Profile) Backend() store.SimpleStore {
	return p.p.getBackend(p)
}

func (p *Profile) Prefix() string {
	return "profiles"
}

func (p *Profile) Key() string {
	return p.Name
}

func (p *Profile) New() store.KeySaver {
	res := &Profile{Name: p.Name, p: p.p}
	return store.KeySaver(res)
}

func (p *Profile) setDT(dp *DataTracker) {
	p.p = dp
}

func (p *Profile) OnCreate() error {
	e := &Error{Code: 409, Type: ValidationError, o: p}
	// We do not allow duplicate profile names
	profiles := AsProfiles(p.p.unlockedFetchAll(p.Prefix()))
	for _, pp := range profiles {
		if pp.Name == p.Name {
			e.Errorf("Profile %s is already exists", p.Name)
			return e
		}
	}
	return nil
}

func (p *Profile) BeforeDelete() error {
	e := &Error{Code: 422, Type: ValidationError, o: p}

	// Make sure no machine is using this profile.
	machines := AsMachines(p.p.unlockedFetchAll(p.p.NewMachine().Prefix()))
	for _, m := range machines {
		if m.HasProfile(p.Name) {
			e.Errorf("Machine %s is using profile %s", m.UUID(), p.Name)
		}
	}

	return e.OrNil()
}

func (p *Profile) OnLoad() error {
	if p.Params == nil {
		p.Params = map[string]interface{}{}
	}
	return nil
}

func (p *Profile) List() []*Profile {
	return AsProfiles(p.p.FetchAll(p))
}

func (p *DataTracker) NewProfile() *Profile {
	return &Profile{p: p, Params: map[string]interface{}{}}
}

func AsProfile(o store.KeySaver) *Profile {
	return o.(*Profile)
}

func AsProfiles(o []store.KeySaver) []*Profile {
	res := make([]*Profile, len(o))
	for i := range o {
		res[i] = AsProfile(o[i])
	}
	return res
}
