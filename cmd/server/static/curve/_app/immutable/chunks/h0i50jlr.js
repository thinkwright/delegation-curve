import{q as n}from"./FRUaXStm.js";const c=/^[a-z0-9-]+$/;function i(e){if(!c.test(e))throw new Error(`Invalid identifier: ${e}`);return e}const o={meta:"SELECT * FROM meta LIMIT 1",delegationAll:`
		SELECT id, name, full_name, score, previous_score,
		       trend, status, weight, tier, description
		FROM delegation
		ORDER BY score DESC
	`,delegationById:e=>`
		SELECT * FROM delegation WHERE id = '${i(e)}'
	`,subIndicators:e=>`
		SELECT name, value, unit, source, freshness
		FROM sub_indicators
		WHERE domain_id = '${i(e)}'
	`,dataSources:e=>`
		SELECT name, cadence, type
		FROM data_sources
		WHERE domain_id = '${i(e)}'
	`};async function l(){const[e]=await n(o.meta);return{delegation:{current:e.delegation_composite,previous:e.delegation_previous,delta:e.delegation_delta,trend:JSON.parse(e.delegation_trend),lastUpdated:e.last_updated,dataYear:e.data_year},domainsTracked:e.domains_tracked,highestDomain:{name:e.highest_domain_name,score:e.highest_domain_score},dataFreshness:e.data_freshness}}async function m(){return(await n(o.delegationAll)).map(a=>({id:a.id,name:a.name,fullName:a.full_name,score:a.score,previousScore:a.previous_score,trend:JSON.parse(a.trend),status:a.status,weight:a.weight,tier:a.tier,subIndicators:[],dataSources:[],description:a.description}))}async function p(e){const[a,r,d]=await Promise.all([n(o.delegationById(e)),n(o.subIndicators(e)),n(o.dataSources(e))]);if(!a.length)return null;const t=a[0];return{id:t.id,name:t.name,fullName:t.full_name,score:t.score,previousScore:t.previous_score,trend:JSON.parse(t.trend),status:t.status,weight:t.weight,tier:t.tier,subIndicators:r.map(s=>({name:s.name,value:s.value,unit:s.unit,source:s.source,freshness:s.freshness})),dataSources:d.map(s=>({name:s.name,cadence:s.cadence,type:s.type})),description:t.description}}export{m as a,p as b,l as g};
