document.body.onload = () => {
  fetch("/devices")
    .then((response) => {
      return response.json();
    })
    .then((json) => {
      console.log(json);
      for (let device of json) {
        let span = document.createElement("span");
        document.body.appendChild(span);
        let img = document.createElement("img");
        img.src = "icons/" + device.Hostname;
        span.appendChild(img);
        let a = document.createElement("a");
        a.href = "http://" + device.Href;
        if (device.Alias) {
          a.innerHTML = device.Alias;
        } else {
          a.innerHTML = device.Hostname;
        }
        span.appendChild(a);
      }
    });
};
