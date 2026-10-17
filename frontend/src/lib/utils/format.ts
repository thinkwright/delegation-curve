export function formatScore(n: number): string {
	return Math.round(n).toString();
}

export function formatDelta(n: number): string {
	const sign = n >= 0 ? '+' : '';
	return `${sign}${n.toFixed(1)}`;
}

export function formatPercent(n: number): string {
	return `${n.toFixed(1)}%`;
}

const MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

/** Formats an ISO date at its own precision: 2025, Sep 2025, or 30 Sep 2025. */
export function formatIsoDate(iso: string): string {
	const [y, m, d] = iso.split('-');
	if (!m) return y;
	const month = MONTHS[Number(m) - 1] ?? m;
	return d ? `${Number(d)} ${month} ${y}` : `${month} ${y}`;
}

/** Formats a measurement window, collapsing a single-period window. */
export function formatWindow(start?: string, end?: string): string {
	if (!start && !end) return '';
	if (!start) return `through ${formatIsoDate(end!)}`;
	if (!end || start === end) return formatIsoDate(start);
	return `${formatIsoDate(start)} – ${formatIsoDate(end)}`;
}
