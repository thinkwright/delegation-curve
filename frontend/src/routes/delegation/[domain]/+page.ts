export function load({ params }: { params: { domain: string } }) {
	return { domainId: params.domain };
}
