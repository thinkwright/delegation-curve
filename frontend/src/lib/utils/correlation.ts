export function getCorrelationColor(value: number): string {
	if (value >= 0.8) return 'bg-black text-white';
	if (value >= 0.6) return 'bg-neutral-700 text-white';
	if (value >= 0.4) return 'bg-neutral-400 text-white';
	if (value >= 0.2) return 'bg-neutral-200 text-black';
	return 'bg-white text-neutral-400';
}
