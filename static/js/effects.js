const techs = document.querySelectorAll("#tech img")

for (let i = 0; i < techs.length; i++) {
  techs[i].addEventListener("mouseover", () => {
    animationDuration = 1;    // second(s)

    techs[i].style.animation = `spinning ${animationDuration}s ease-in-out`;
    setTimeout(() => {
      techs[i].style.animation = "none";
    }, animationDuration * 1000)
  })
}
