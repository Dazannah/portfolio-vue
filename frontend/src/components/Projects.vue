<template>
  <div id="projects" class="hero bg-base-200 scroll-mt-16">
    <div class="hero-content">
      <h1 class="text-3xl font-bold">Projektek</h1>
    </div>
  </div>
  <div class="max-w-screen bg-base-200 py-6">
    <div class="max-w-8/10 m-auto flex flex-col lg:flex-row gap-2 items-stretch">
      <div class="card bg-base-100 w-96 h-full shadow-sm flex-auto">
        <figure>
          <img src="../assets/imgs/car-register.png" />
        </figure>
        <div class="card-body">
          <h2 class="card-title">Gépjármű nyilvántartás</h2>
          <p>
            Felhasználók kezelése admin felületen. Gépjármű gyors felvétele, gyors leadása,
            előfoglalása.
          </p>
          <p class="pt-1">Laravel, Tailwind CSS, Blade, Livewire, MySQL</p>
          <div class="card-actions justify-end">
            <a href="" target="_blank"><button class="btn btn-primary">Demó</button></a>
            <a href="" target="_blank"><button class="btn btn-primary">Github</button></a>
          </div>
        </div>
      </div>

      <div class="card bg-base-100 w-96 h-full shadow-sm flex-auto">
        <figure>
          <img src="../assets/imgs/usermanager.png" />
        </figure>
        <div class="card-body">
          <h2 class="card-title">Felhasználó kezelő</h2>
          <p>
            Felhasználók, igénylése, szerkesztése, törlése. Hozzájuk tartozó jogosultságok kezelése.
          </p>
          <p class="pt-1">Laravel, Tailwind CSS, Blade, Livewire, MySQL</p>
          <div class="card-actions justify-end">
            <a href="" target="_blank"><button class="btn btn-primary">Demó</button></a>
            <a href="" target="_blank"> <button class="btn btn-primary">Github</button></a>
          </div>
        </div>
      </div>

      <div class="card bg-base-100 w-96 h-full shadow-sm flex-auto">
        <figure>
          <img src="../assets/imgs/appointment.png" />
        </figure>
        <div class="card-body">
          <h2 class="card-title">Appointment</h2>
          <p>Időpontok foglalása, kezelése. Felhasználók kezelése. Szünnapok megadása.</p>
          <p class="pt-1">Laravel, Tailwind CSS, Blade, MySQL</p>
          <div class="card-actions justify-end">
            <a href="" target="_blank"> <button class="btn btn-primary">Demó</button></a>
            <a href="" target="_blank"> <button class="btn btn-primary">Github</button></a>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
