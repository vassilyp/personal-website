const techs = document.querySelectorAll("#tech img")
const animationDuration = 1;    // second(s)

for (let i = 0; i < techs.length; i++) {
  techs[i].addEventListener("mouseover", () => {

    if (techs[i].style.animation === "") {
      techs[i].style.animation = `spinning ${animationDuration}s ease-in-out`;
      setTimeout(() => {
        techs[i].style.animation = "";
      }, animationDuration * 1000)
    }
  })
}
