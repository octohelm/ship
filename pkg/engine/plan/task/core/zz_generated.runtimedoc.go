/*
Package core GENERATED BY gengo:runtimedoc 
DON'T EDIT THIS FILE
*/
package core

// nolint:deadcode,unused
func runtimeDoc(v any, names ...string) ([]string, bool) {
	if c, ok := v.(interface {
		RuntimeDoc(names ...string) ([]string, bool)
	}); ok {
		return c.RuntimeDoc(names...)
	}
	return nil, false
}

func (v Auth) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Username":
			return []string{}, true
		case "Secret":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v CacheDir) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "ID":
			return []string{}, true
		case "Concurrency":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Export) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Exporter":
			return []string{}, true

		}
		if doc, ok := runtimeDoc(v.Exporter, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v FS) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {

		}

		return nil, false
	}
	return []string{}, true
}

func (v Image) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Rootfs":
			return []string{}, true
		case "Config":
			return []string{}, true
		case "Platform":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v ImageConfig) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "WorkingDir":
			return []string{}, true
		case "Env":
			return []string{}, true
		case "Labels":
			return []string{}, true
		case "Entrypoint":
			return []string{}, true
		case "Cmd":
			return []string{}, true
		case "User":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Mount) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Mounter":
			return []string{}, true

		}
		if doc, ok := runtimeDoc(v.Mounter, names...); ok {
			return doc, ok
		}

		return nil, false
	}
	return []string{}, true
}

func (v MountCacheDir) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Type":
			return []string{}, true
		case "Dest":
			return []string{}, true
		case "Contents":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v MountFile) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Type":
			return []string{}, true
		case "Dest":
			return []string{}, true
		case "Contents":
			return []string{}, true
		case "Permissions":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v MountFs) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Type":
			return []string{}, true
		case "Dest":
			return []string{}, true
		case "Contents":
			return []string{}, true
		case "Source":
			return []string{}, true
		case "ReadOnly":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v MountSecret) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Type":
			return []string{}, true
		case "Dest":
			return []string{}, true
		case "Contents":
			return []string{}, true
		case "Uid":
			return []string{}, true
		case "Gid":
			return []string{}, true
		case "Mask":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v MountTemp) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Type":
			return []string{}, true
		case "Dest":
			return []string{}, true
		case "Contents":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v Secret) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {

		}

		return nil, false
	}
	return []string{}, true
}

func (v SecretOrString) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Value":
			return []string{}, true
		case "Secret":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v StringOrBool) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "String":
			return []string{}, true
		case "Bool":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}

func (v TempDir) RuntimeDoc(names ...string) ([]string, bool) {
	if len(names) > 0 {
		switch names[0] {
		case "Size":
			return []string{}, true

		}

		return nil, false
	}
	return []string{}, true
}
