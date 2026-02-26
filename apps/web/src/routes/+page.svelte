<script lang="ts">
  // Состояние для Svelte 5
  let weather = $state({ temp: '--', city: 'Moscow', loading: true });

  async function getWeatherData() {
    try {
      // Стучимся в наш Go-прокси через Vercel/Vite rewrite
      const res = await fetch('/api/index');
      const data = await res.json();
      weather.temp = Math.round(data.main.temp);
      weather.loading = false;
    } catch (e) {
      console.error("Go backend is sleeping or blocked...");
      weather.temp = "!?";
    }
  }

  // Авто-запуск при загрузке
  $effect(() => { getWeatherData(); });
</script>

<main class="min-h-screen flex items-center justify-center bg-[#050505] font-sans selection:bg-cyan-500/30">
  <div class="relative group">
    <!-- Тот самый дизайнерский глоу -->
    <div class="absolute -inset-1 bg-gradient-to-tr from-cyan-500 to-blue-600 rounded-[2.5rem] blur opacity-20 group-hover:opacity-40 transition duration-1000"></div>
    
    <div class="relative bg-black/40 border border-white/10 backdrop-blur-3xl p-12 rounded-[2.5rem] w-[350px] shadow-2xl">
      <header class="flex justify-between items-start mb-12">
        <div>
          <h1 class="text-white/30 text-[10px] uppercase tracking-[0.4em] font-bold mb-1">Station 01</h1>
          <p class="text-white/90 text-sm font-light tracking-widest">{weather.city}</p>
        </div>
        <div class="w-2 h-2 rounded-full bg-cyan-500 shadow-[0_0_15px_rgba(6,182,212,0.5)] animate-pulse"></div>
      </header>

      <div class="mb-12">
        <span class="text-white text-8xl font-extralight tracking-tighter leading-none select-none">
          {weather.temp}<span class="text-cyan-500 font-normal">°</span>
        </span>
      </div>

      <footer class="flex flex-col items-center justify-center pt-8 border-t border-white/5 gap-2">
    <!-- The "Love" Note -->
    <div class="flex items-center gap-1.5 group/love cursor-default">
        <span class="text-[10px] text-white/20 uppercase tracking-[0.2em] font-medium group-hover/love:text-white/40 transition-colors">
            Made with
        </span>
        <!-- Pulsing Heart -->
        <span class="text-rose-500 animate-pulse text-xs">❤️</span>
        <span class="text-[10px] text-white/20 uppercase tracking-[0.2em] font-medium group-hover/love:text-white/40 transition-colors">
            for you
        </span>
    </div>

    <!-- Tech Stack (Small & Subtle) -->
    <div class="flex items-center gap-3 opacity-20 hover:opacity-50 transition-opacity duration-500">
        <span class="text-[8px] font-mono tracking-tighter text-cyan-400">Go v1.22</span>
        <div class="w-1 h-1 rounded-full bg-white/20"></div>
        <span class="text-[8px] font-mono tracking-tighter text-orange-400">Svelte 5</span>
    </div>
</footer>

    </div>
  </div>
</main>