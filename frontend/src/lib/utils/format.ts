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
