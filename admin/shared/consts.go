package shared

import "github.com/dracory/cdn"

const PathHome = "home"
const PathGroups = "groups"
const PathRoles = "roles"
const PathUsers = "users"
const PathUserCreate = "user-create"
const PathUserUpdate = "user-update"
const PathUserDelete = "user-delete"
const PathUserImpersonate = "user-impersonate"

var ScriptHtmx = `setTimeout(async function() {
	if (!window.htmx) {
		let script = document.createElement('script');
		document.head.appendChild(script);
		script.type = 'text/javascript';
		script.src = '` + cdn.Htmx_2_0_0() + `';
		await script.onload
	}
}, 1000);`

var ScriptSwal = `setTimeout(async function() {
	if (!window.Swal) {
		let script = document.createElement('script');
		document.head.appendChild(script);
		script.type = 'text/javascript';
		script.src = '` + cdn.Sweetalert2_11() + `';
		await script.onload
	}
}, 1000);`
