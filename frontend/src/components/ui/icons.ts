// Inner SVG markup for a 24x24 viewBox. Stroke/fill defaults live in BaseIcon.
export const icons = {
  play: '<path d="M7 5l12 7-12 7z"/>',
  stop: '<rect x="6" y="6" width="12" height="12" rx="1.5"/>',
  pause: '<path d="M9 6v12M15 6v12"/>',
  restart: '<path d="M20 12a8 8 0 1 1-2.3-5.7"/><path d="M20 4v4h-4"/>',
  trash: '<path d="M4 7h16M9 7V4h6v3M6 7l1 13h10l1-13"/>',
  box: '<rect x="3" y="7" width="18" height="12" rx="1.5"/><path d="M7 7v12M11 7v12M15 7v12M3 7l3-3h12l3 3"/>',
  layers: '<path d="M12 4l9 5-9 5-9-5z"/><path d="M3 14l9 5 9-5"/>',
  disk: '<ellipse cx="12" cy="6" rx="8" ry="3"/><path d="M4 6v12c0 1.7 3.6 3 8 3s8-1.3 8-3V6"/><path d="M4 12c0 1.7 3.6 3 8 3s8-1.3 8-3"/>',
  net: '<circle cx="12" cy="5" r="2"/><circle cx="5" cy="19" r="2"/><circle cx="19" cy="19" r="2"/><path d="M12 7v5M12 12l-6 5M12 12l6 5"/>',
  down: '<path d="M12 4v11M7 10l5 5 5-5M5 20h14"/>',
  plus: '<path d="M12 5v14M5 12h14"/>',
  theme: '<circle cx="12" cy="12" r="8"/><path d="M12 4a8 8 0 0 0 0 16z" fill="currentColor"/>',
} as const

export type IconName = keyof typeof icons
