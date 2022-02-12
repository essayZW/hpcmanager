package verify

import "testing"

var defaultVerify Verify

func init() {
	defaultVerify = &hardcodeVerify{
		actionsLevel: map[PermissionAction]*actionVerify{
			"ACTION1": {
				maxLevel: SuperAdmin,
				minLevel: SuperAdmin,
			},
			"ACTION2": {
				maxLevel: SuperAdmin,
				minLevel: CommonAdmin,
			},
			"ACTION3": {
				maxLevel: MaxLevel,
				minLevel: Common,
			},
			"ACTION4": {
				maxLevel: Tutor,
				minLevel: MinLevel,
			},
		},
	}
}

func TestDefaultVerifyIdentify(t *testing.T) {

	tests := []struct {
		name   string
		action PermissionAction
		levels []Level

		except bool
	}{
		{
			name:   "single permission forbidden",
			action: "ACTION1",
			levels: []Level{
				Common,
			},
			except: false,
		},
		{
			name:   "single permission accept",
			action: "ACTION1",
			levels: []Level{
				SuperAdmin,
			},
			except: true,
		},
		{
			name:   "single permission accept2",
			action: "ACTION1",
			levels: []Level{
				SuperAdmin,
				Tutor,
			},
			except: true,
		},
		{
			name:   "range permission forbidden",
			action: "ACTION2",
			levels: []Level{
				Guest,
			},
			except: false,
		},
		{
			name:   "range permission accept",
			action: "ACTION2",
			levels: []Level{
				SuperAdmin,
				Tutor,
			},
			except: true,
		},
		{
			name:   "large than Common forbidden",
			action: "ACTION3",
			levels: []Level{
				Guest,
			},
			except: false,
		},
		{
			name:   "large than Common accept",
			action: "ACTION3",
			levels: []Level{
				Tutor,
			},
			except: true,
		},
		{
			name:   "less than Tutor forbidden",
			action: "ACTION4",
			levels: []Level{
				CommonAdmin,
			},
			except: false,
		},
		{
			name:   "less than Tutor accept",
			action: "ACTION4",
			levels: []Level{
				Common,
			},
			except: true,
		},
		{
			name:   "unsupproted action",
			action: "NO",
			levels: []Level{
				SuperAdmin,
			},
			except: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := defaultVerify.Identify(test.action, test.levels)
			if res != test.except {
				t.Errorf("Except %v Get %v", test.except, res)
			}
		})
	}
}
