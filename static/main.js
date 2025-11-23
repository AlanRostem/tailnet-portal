document.body.onload = () => {
  fetch("/devices")
    .then((response) => {
      return response.json();
    })
    .then((json) => {
      console.log(json);
      let container = document.getElementById("container");
      for (let device of json) {
        let a = document.createElement("a");
        a.className = "link-span";
        a.href = "http://" + device.Href;
        container.appendChild(a);
        let img = document.createElement("img");
        img.src = "icons/" + device.Hostname + ".png";
        a.appendChild(img);
        let p = document.createElement("p");
        if (device.Alias) {
          p.innerHTML = device.Alias;
        } else {
          p.innerHTML = device.Hostname;
        }
        a.appendChild(p);
      }
    });
};
