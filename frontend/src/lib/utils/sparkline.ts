export function generatePath(
	data: number[],
	width: number = 100,
	height: number = 24,
	padding: number = 2,
	yMin?: number,
	yMax?: number
): string {
	if (data.length < 2) return '';
	const min = yMin ?? Math.min(...data);
	const max = yMax ?? Math.max(...data);
	const range = max - min || 1;
	const stepX = width / (data.length - 1);

	return data
		.map((val, i) => {
			const x = i * stepX;
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
