<template>
  <div id="hero" class="lg:hero p-5 bg-base-200 min-h-screen max-w-full scroll-mt-16">
    <div class="lg:hero-content flex-col lg:flex-row-reverse max-w-full">
      <img
        src="../assets/imgs/profilePic.jpg"
        class="max-w-2xs mx-auto sm:max-w-xs rounded-lg shadow-2xl"
      />
      <div>
        <h1 class="text-3xl lg:text-5xl font-bold">Üdvözöllek a portfólió weboldalamon!</h1>
        <p class="py-6">
          Fábián Dávid vagyok, webfejlesztő. 2023 márciusában kezdtem a Junior Fullstack API
          fejlesztő képzést az IEC-nél. Ez a képzés mellett egy pár kurzust az Udemy-n is elvégeztem
          különböző témákban. A teljesség igénye nélkül:
        </p>
        <p class="py-6">Laravel PHPUnit NodeJS ReactJS</p>
        <p class="py-6">
          Szeretem a már valamennyire ismert, használt technológiákról szerzett tudásomat
          fejleszteni, és szeretek új technológiákat kipróbálni, alkalmazni. Főleg a backend
          érdekel, ezen a területen szerzett tudásomat szeretném elsősorban elmélyíteni, de azért a
          frontend-el is elboldogulok. Alább letölthető az önéletrajzom, illetve megtekinthető
          néhány gyakorló/tanuló projektem.
        </p>

        <a href="/cv.pdf" target="_blank"><button class="btn btn-primary">Önéletrajz</button></a>
        <a href="#projects"><button class="btn">Projektek</button></a>
      </div>
    </div>
  </div>
</template>
