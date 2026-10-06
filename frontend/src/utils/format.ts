const rtf = new Intl.RelativeTimeFormat('es', {numeric: 'auto'})

const units: [Intl.RelativeTimeFormatUnit, number][] = [
  ['year', 31536000], ['month', 2592000], ['week', 604800],
  ['day', 86400], ['hour', 3600], ['minute', 60],
]

// "hace 3 horas" from a unix timestamp in seconds.
export function timeAgo(unixSeconds: number, now = Date.now() / 1000): string {
  const diff = unixSeconds - now
  for (const [unit, secs] of units) {
    if (Math.abs(diff) >= secs) return rtf.format(Math.round(diff / secs), unit)
  }
  return rtf.format(0, 'second')
}

// Decimal units, like the Docker CLI ("243 MB").
export function formatBytes(bytes: number): string {
  if (bytes < 0) return '—'
  const units = ['B', 'kB', 'MB', 'GB', 'TB']
  let i = 0
  let n = bytes
  while (n >= 1000 && i < units.length - 1) {
    n /= 1000
    i++
  }
  return `${n < 10 && i > 0 ? n.toFixed(1) : Math.round(n)} ${units[i]}`
}
