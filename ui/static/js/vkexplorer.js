function openURL(event) {
    event.preventDefault();
    let appId = document.getElementById("app_id").value;
    console.log("openURL: " + appId);
    let appUrl = "https://oauth.vk.com/authorize?client_id="
        + appId + "&display=page&redirect_uri=https://oauth.vk.com/blank.html&scope=notify,friends,photos,audio,video,docs,notes,pages,status,wall,groups,notifications&response_type=token&v=5.131"
    // open in new page
    window.open(appUrl, "_blank");
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-app-id", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            location.reload()
        }
    };

    xhr.send("app_id=" + appId);
}

function postURL(event) {
    event.preventDefault();
    let appUrl = document.getElementById("app_url").value;
    console.log("postURL: " + appUrl);

    // window.open("https://google.com", "_blank");
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/set-app-url", true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            location.reload()
        }
    };
    // send
    let prm = "app_url=" + appUrl;
    console.log("send to: " + prm);
    xhr.send(prm);
}

function openTestURL(ibase) {
    let base = document.getElementById(ibase).textContent;
    console.log(base)
    let token = document.getElementById('token').textContent;
    console.log(token)
    let url = base + token;
    window.open(url, "_blank");
}

// HELPERS FOR UPDATE DB QUERIES

function sendUpdateQuery(url) {
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    // handler
    xhr.onreadystatechange = function() {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            //location.reload()
            // while the handler function is unnecessary?
        }
    };

    console.log(url);
    xhr.send();
}

function sendUpdateQueryARG(url, argname, argvalue) {
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

    const data = argname + "=" + argvalue;

    console.log(url);
    xhr.send(data);
}

function sendUpdateQueryCB(url) {
    // prepare AJAX
    let xhr = new XMLHttpRequest();
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");

    const checkboxes = document.querySelectorAll("input[type=checkbox]:checked");
    const checkboxValues = [];
    for (let i = 0; i < checkboxes.length; i++) {
        if(checkboxes[i].id != "all")
            checkboxValues.push(checkboxes[i].id);
    }
    const data = "checkbox=" + encodeURIComponent(checkboxValues.join(","));

    xhr.send(data);
    console.log(url);
 }
 
// UPDATE DB QUERIES

function updateMyFriends() {
    sendUpdateQuery("/update-my-friends")
}

function updateMyGroups() {
    sendUpdateQuery("/update-my-groups")
}

function updateMyBookmarks() {
    sendUpdateQuery("/update-my-bookmarks")
}

function updateCheckedGroupMembers() {
    sendUpdateQueryCB("/update-checked-group-members")
}

function updateCheckedGroupWall() {
    sendUpdateQueryCB("/update-checked-group-wall")
}

function updateCheckedUserFriends() {
    sendUpdateQueryCB("/update-checked-user-friends")
}

function updateCheckedUserWall() {
    sendUpdateQueryCB("/update-checked-user-wall")
}

function updateCheckedUserGroups() {
    sendUpdateQueryCB("/update-checked-user-groups")
}

function updateUserInfo() {
    let id = document.getElementById("user_id").value;
    sendUpdateQueryARG("/update-user-info", "user", id)
}

function updateUserFriends() {
	let id = document.getElementById("user_id").value;
	sendUpdateQueryARG("/update-user-friends", "user", id)
}

function updateUserGroups() {
	let id = document.getElementById("user_id").value;
	sendUpdateQueryARG("/update-user-groups", "user", id)
}

function updateUserWall() {
	let id = document.getElementById("user_id").value;
	sendUpdateQueryARG("/updateuserwall", "user", id)
}

function updateGroupInfo() {
    let id = document.getElementById("group_id").value;
    sendUpdateQueryARG("/update-group-info", "group", id)
}

function updateGroupMembers() {
    let id = document.getElementById("group_id").value;
    sendUpdateQueryARG("/update-group-members", "group", id)
}

function updateGroupWall() {
    let id = document.getElementById("group_id").value;
    sendUpdateQueryARG("/update-group-wall", "group", id)
}

// GLOBAL AREA

document.addEventListener("DOMContentLoaded", function(event) {
    console.log("init page")

    let statusbar = document.getElementById('statusbar');

    // Update the status in the HTML
    let source = new EventSource("http://localhost:8080/get-server-status");
    source.onmessage = function(event) {
        statusbar.textContent = event.data;
    };

    //
    let started = false;
    let btn = document.getElementById("controlbtn");

    // Get the initial state of the goroutine
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/get-worker-status");
    xhr.onload = function() {
        if (xhr.status === 200) {
            started = (xhr.responseText === "true");
            btn.textContent = started ? "STOP" : "START";
        }
    };
    xhr.send();

    btn.addEventListener("click", function(event) {
        event.preventDefault();
        if (started) {
            console.log("stop...")
            btn.textContent = "STOPPING...";
            // Stop the goroutine
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://127.0.0.1:8080/stop-worker");
            xhr.onload = function() {
                console.log("Worker stopped.");
                started = false;
                btn.textContent = "START";
            };
            xhr.onerror = function() {
                console.log("Error stopping worker.");
            };
            xhr.send();
        } else {
            console.log("start...")
            btn.textContent = "STARTING...";
            // Start the goroutine
            let xhr = new XMLHttpRequest();
            xhr.open("POST", "http://127.0.0.1:8080/start-worker");
            xhr.onload = function() {
                console.log("Worker started.");
                started = true;
                btn.textContent = "STOP";
            };
            xhr.onerror = function() {
                console.log("Error starting worker.");
            };
            xhr.send();
        }
    });
});
