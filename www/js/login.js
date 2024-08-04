document.addEventListener("DOMContentLoaded", function() {
    const loginForm = document.getElementById("login-form");
    loginForm.addEventListener("submit", function(event) {
      event.preventDefault();
      
      const formData = new FormData(loginForm);
      const jsonData = {};

      formData.forEach((value, key) => {
        jsonData[key] = value;
      });
 
      fetch("/api/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(jsonData)
      })
      .then(loginForm.reset())
      .then(response => {
        if(response.ok) {
          console.log("login successful")
          var expires = (new Date(Date.now()+ 86400*1000)).toUTCString();
          document.cookie = "X-Auth-Token=" + response.headers.get("X-Auth-Token") + "; expires=" + expires + "; path=/"
          window.location = "/admin-panel"
        }
        else if(response.status === 403) {
          window.alert("invalid username or password")
        } else {
          console.error(response.text)
        }
      })
      .catch(error => {
        console.error("There was a problem with the fetch operation:", error);
      });
    });
  });
  