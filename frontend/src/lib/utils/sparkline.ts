/**
 * Relative x positions (0..1) for a series. With dates (YYYY-MM-DD), points
 * are placed proportionally in time so uneven gaps between runs keep their
 * true width; otherwise points are evenly spaced.
 */
export function relativeX(count: number, dates?: (string | undefined)[]): number[] {
	if (count <= 1) return Array.from({ length: count }, () => 1);
	const times = dates?.length === count ? dates.map((d) => (d ? Date.parse(d) : NaN)) : [];
	const valid = times.length === count && times.every((t) => Number.isFinite(t));
	const first = valid ? times[0] : 0;
	const span = valid ? times[count - 1] - first : 0;
	if (!valid || span <= 0) return Array.from({ length: count }, (_, i) => i / (count - 1));
	return times.map((t) => (t - first) / span);
}

export function generatePath(
	data: number[],
	width: number = 100,
	height: number = 24,
	padding: number = 2,
	yMin?: number,
	yMax?: number,
	dates?: (string | undefined)[]
): string {
	if (data.length < 2) return '';
	const min = yMin ?? Math.min(...data);
	const max = yMax ?? Math.max(...data);
	const range = max - min || 1;
	const xs = relativeX(data.length, dates);

	return data
		.map((val, i) => {
			const x = xs[i] * width;
			const y = height - padding - ((val - min) / range) * (height - padding * 2);
			return `${i === 0 ? 'M' : 'L'}${x.toFixed(1)},${y.toFixed(1)}`;
		})
		.join(' ');
}

export function generateAreaPath(
	data: number[],
	width: number = 100,
	height: number = 24,
	padding: number = 2,
	yMin?: number,
	yMax?: number
): string {
	const linePath = generatePath(data, width, height, padding, yMin, yMax);
	if (!linePath) return '';
	return `${linePath} L${width.toFixed(1)},${height} L0,${height} Z`;
}
