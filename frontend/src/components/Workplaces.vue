<template>
  <div id="workplaces" class="hero bg-base-200 lg:scroll-mt-16">
    <div class="hero-content">
      <h1 class="text-3xl font-bold">Munkahelyek</h1>
    </div>
  </div>
  <div class="hero bg-base-200">
    <div class="hero-content">
      <ul class="timeline timeline-snap-icon max-md:timeline-compact timeline-vertical">
        <li>
          <div class="timeline-middle">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
              class="h-5 w-5"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div class="timeline-start mb-10 md:text-end">
            <time class="font-mono">2020 február - jelenleg is</time>
            <div class="text-lg font-black">Infolan Kft.</div>
            <div class="text-md">Rendszergazda</div>
            <p>
              Rendszergazdai feladatok elvégzése mellett, CSMEKHM intrawebes jogosultságigénylő
              alkalmazás fejlesztése, karbantartása. NodeJS, MongoDB, és React-ot használva.
              Jelenleg egy általánosabb, testreszabható verzió fejlesztésén dolgozom, Laravel,
              Blade, Livewire és MySQL-el, amely a kórházi rendszerből származó igényekre épül.
            </p>
            <p class="pt-2">Hódmezővásárhely, Makó</p>
          </div>
          <hr />
        </li>
        <li>
          <hr />
          <div class="timeline-middle">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 20 20"
              fill="currentColor"
              class="h-5 w-5"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.857-9.809a.75.75 0 00-1.214-.882l-3.483 4.79-1.88-1.88a.75.75 0 10-1.06 1.061l2.5 2.5a.75.75 0 001.137-.089l4-5.5z"
                clip-rule="evenodd"
              />
            </svg>
          </div>
          <div class="timeline-end md:mb-10">
            <time class="font-mono">2019 január - 2020 január</time>
            <div class="text-lg font-black">Csongrád megyei Pedagógiai Szakszolgálat</div>
            <div class="text-md">Rendszergazda</div>
            <p class="pt-2">Szeged, Makó, Hódmezővásárhely</p>
          </div>
          <hr />
        </li>
      </ul>
    </div>
  </div>
</template>
