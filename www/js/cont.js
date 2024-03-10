document.addEventListener("DOMContentLoaded", function() {
    const contactForm = document.getElementById("contact-form");
  
    contactForm.addEventListener("submit", function(event) {
      event.preventDefault();
      
      const formData = new FormData(contactForm);
      const jsonData = {};

      console.log(formData);

      formData.forEach((value, key) => {
        jsonData[key] = value;
      });

      console.log(jsonData);
      fetch("/api/contact", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(jsonData)
      })
      .then(response => {
        if(!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then(contactForm.reset())
      .catch(error => {
        console.error("There was a problem with the fetch operation:", error);
        //alert("Oops! Something went wrong. Please try again later.");
      });
    });
  });
  