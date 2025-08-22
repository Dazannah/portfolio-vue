<template>
  <div
    class="navbar bg-primary text-primary-content sticky top-0 z-50 shadow-sm h-16 hidden lg:flex"
  >
    <div class="flex-1">
      <div class="menu menu-horizontal px-1">
        <a class="btn btn-ghost" href="#hero">Üdvözöllek</a>
        <a class="btn btn-ghost" href="#introduction">Bemutatkozás</a>
        <a class="btn btn-ghost" href="#workplaces">Munkahelyek</a>
        <a class="btn btn-ghost" href="#studies">Tanulmányok</a>
        <a class="btn btn-ghost" href="#projects">Projektek</a>
        <a class="btn btn-ghost" href="#footer">Elérhetőségek</a>
      </div>
    </div>

    <div class="flex-none">
      <div class="menu menu-horizontal px-1">
        <a class="btn btn-ghost" href="/cv.pdf" target="_blank">Önéletrajz</a>
      </div>
    </div>
  </div>

  <div class="navbar bg-primary sticky top-0 z-50 shadow-sm h-16 shadow-sm lg:hidden block">
    <div class="navbar-start">
      <div class="dropdown">
        <div tabindex="0" role="button" class="btn btn-ghost btn-circle text-primary-content">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-5 w-5"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M4 6h16M4 12h16M4 18h7"
            />
          </svg>
        </div>
        <ul
          tabindex="0"
          class="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow"
        >
          <a class="btn btn-ghost" href="#hero">Üdvözöllek</a>
          <a class="btn btn-ghost" href="#introduction">Bemutatkozás</a>
          <a class="btn btn-ghost" href="#workplaces">Munkahelyek</a>
          <a class="btn btn-ghost" href="#studies">Tanulmányok</a>
          <a class="btn btn-ghost" href="#projects">Projektek</a>
          <a class="btn btn-ghost" href="#footer">Elérhetőségek</a>
        </ul>
      </div>
    </div>
    <div class="navbar-center"></div>
    <div class="navbar-end">
      <a class="btn btn-ghost text-primary-content" href="/cv.pdf" target="_blank">Önéletrajz</a>
    </div>
  </div>
</template>
