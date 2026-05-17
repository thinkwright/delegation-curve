import type { SourceDisclosure } from './types';

export const sourceDisclosures: Record<string, SourceDisclosure[]> = {
	'content-mod': [
		{
			name: 'Meta 2025 Enforcement Policy Update',
			role: 'Methodology-break context',
			cadence: 'Periodic',
			type: 'policy update',
			url: 'https://about.fb.com/news/2025/01/meta-more-speech-fewer-mistakes/amp/',
			note: 'Documents Meta policy changes that affect comparability of proactive enforcement rates; the scored Meta value is kept with a break caveat rather than treated as a simple trend continuation.'
		},
		{
			name: 'YouTube Transparency Report Method Notes',
			role: 'Source-semantics guardrail',
			cadence: 'Quarterly',
			type: 'documentation',
			url: 'https://support.google.com/transparencyreport/answer/9198203?hl=en-GB',
			note: 'Supports the 99.5% automated flagging input and clarifies that this indicator measures automated detection/flagging rather than final human-reviewed removal.'
		},
		{
			name: 'TikTok Community Guidelines Reports',
			role: 'Denominator context',
			cadence: 'Quarterly',
			type: 'transparency-report',
			url: 'https://www.tiktok.com/transparency/en-us/reports/',
			note: 'Used to compare EU DSA automated enforcement against global proactive-detection reporting; the current score uses the DSA without-human-review metric.'
		},
		{
			name: 'X Global and DSA Transparency Reports',
			role: 'Held platform context',
			cadence: 'Semi-annual',
			type: 'transparency-report',
			url: 'https://transparency.x.com/content/dam/transparency-twitter/2025/x-global-transparency-report_h2_2024.pdf',
			note: 'Retained as context because comparable 2025/2026 global automation rates were not locked; category-level DSA rates vary too sharply to score directly.'
		},
		{
			name: 'EU DSA Transparency Database Research API',
			role: 'Future extraction candidate',
			cadence: 'Continuous',
			type: 'api',
			url: 'https://transparency.dsa.ec.europa.eu/page/research-api',
			note: 'A promising cross-platform automation source, but not scored until API access, query definitions, and run-level snapshots are reproducible.'
		}
	],
	'algo-trade': [
		{
			name: 'BIS 2025 FX Execution Analysis',
			role: 'Scored electronic-execution input',
			cadence: 'Triennial',
			type: 'survey',
			url: 'https://www.bis.org/publ/qtrpdf/r_qt2512v.htm',
			note: 'Provides the 59% global FX electronic trading share. It is the cleanest refreshable trading input, but electronic execution is not equivalent to AI-directed trading.'
		},
		{
			name: 'Coalition Greenwich Buy-Side AI Trading Survey',
			role: 'Scored current-AI input',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.greenwich.com/market-structure-technology/great-expectations-ai-equity-trading',
			note: 'Uses the 15% current internal-AI trade-execution value; higher current-plus-planned and expected-impact figures are treated as context rather than current adoption.'
		},
		{
			name: 'SEC Market Structure Analytics',
			role: 'Market-structure context',
			cadence: 'Periodic',
			type: 'regulatory data',
			url: 'https://www.sec.gov/featured-topics/market-structure-analytics',
			note: 'Authoritative automation and market-structure context, but not used as a direct AI delegation source without a clearer algorithmic or AI execution denominator.'
		},
		{
			name: 'Cboe US Equities Year in Review',
			role: 'Electronic-market context',
			cadence: 'Annual',
			type: 'market report',
			url: 'https://www.cboe.com/insights/posts/2025-u-s-equities-year-in-review/',
			note: 'Supports the scale of electronic market structure while remaining too broad to score as AI or algorithmic trading volume.'
		},
		{
			name: 'The TRADE Algorithmic Trading Survey 2025',
			role: 'Adoption context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.thetradenews.com/wp-content/uploads/2025/04/Algo56-FINAL.pdf',
			note: 'Useful for practitioner adoption context, but not scored because the response base and denominator do not map cleanly to market-wide delegated execution share.'
		},
		{
			name: 'Cboe Options Industry Reports',
			role: 'Demoted context',
			cadence: 'Annual',
			type: 'market report',
			url: 'https://www.cboe.com/insights/posts/the-state-of-the-options-industry-2025/',
			note: 'Fresh options-market volume source, but it does not directly measure automated or AI-driven options execution share.'
		}
	],
	'code-gen': [
		{
			name: 'Sonar State of Code 2026',
			role: 'Scored output-share component',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.sonarsource.com/blog/state-of-code-developer-survey-report-the-current-reality-of-ai-coding',
			note: 'Contributes the 42% committed-code estimate inside the blended AI-generated output share; treated as survey evidence, not repository telemetry.'
		},
		{
			name: 'GitLab Global DevSecOps Report 2026',
			role: 'Scored output-share triangulation',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://about.gitlab.com/resources/developer-survey/',
			note: 'Provides the 34% AI-generated code-source-share comparator used to temper the higher Sonar and AI-forward estimates.'
		},
		{
			name: 'METR AI Usage Survey',
			role: 'Scored value-share input',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://metr.org/blog/2026-05-11-ai-usage-survey/',
			note: 'Uses the median March 2026 2x self-reported value multiplier as a 50% AI-attributable technical-work value-share estimate.'
		},
		{
			name: 'Google DORA 2025 AI-Assisted Software Development',
			role: 'Scored workflow-reliance component',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://blog.google/innovation-and-ai/technology/developers-tools/dora-report-2025/',
			note: 'Contributes the moderate-to-heavy reliance signal because it measures work time and reliance, not just whether developers have tried AI tools.'
		},
		{
			name: 'JetBrains AI Pulse 2026',
			role: 'Scored workflow-reliance triangulation',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://blog.jetbrains.com/research/2026/04/which-ai-coding-tools-do-developers-actually-use-at-work/',
			note: 'Triangulates regular work use and specialized developer-tool adoption, separating developer agents/editors from general chatbots.'
		},
		{
			name: 'Anthropic Agentic Coding Trends Report',
			role: 'Scored delegation calibration',
			cadence: 'Annual',
			type: 'report',
			url: 'https://resources.anthropic.com/hubfs/2026%20Agentic%20Coding%20Trends%20Report.pdf?hsLang=en',
			note: 'Used for the assisted-versus-fully-delegated task split, aligning the score with delegation rather than general AI involvement.'
		},
		{
			name: 'Agentic Much? Coding-Agent Adoption Study',
			role: 'Scored behavioral adoption component',
			cadence: 'Periodic',
			type: 'research',
			url: 'https://arxiv.org/abs/2601.18341',
			note: 'Open-source GitHub trace evidence helps ground agentic adoption behaviorally, while remaining narrower than private enterprise software work.'
		},
		{
			name: 'GitHub Octoverse 2025',
			role: 'Platform context',
			cadence: 'Annual',
			type: 'report',
			url: 'https://github.blog/news-insights/octoverse/octoverse-a-new-developer-joins-github-every-second-as-ai-leads-typescript-to-1/',
			note: 'Strong evidence that AI is embedded in developer workflows, but not used as a direct score input because accepted-lines or agent-authored-PR denominators were not locked.'
		},
		{
			name: 'Augment State of AI-Native Engineering 2026',
			role: 'Frontier-adopter context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.augmentcode.com/blog/ai-native-survey-2026',
			note: 'Retained as an AI-forward upper context value, not as a population estimate, because the sample is intentionally concentrated among AI-native engineering teams.'
		},
		{
			name: 'Alphabet Q3 2025 Earnings Call',
			role: 'Frontier-company benchmark',
			cadence: 'Quarterly',
			type: 'company disclosure',
			url: 'https://abc.xyz/investor/events/event-details/2025/2025-Q3-Earnings-Call-2025-4OI4Bac_Q9/default.aspx',
			note: 'Nearly-half AI-generated code benchmark is treated as a top-tech comparison point rather than a broad developer-population input.'
		},
		{
			name: 'Harness State of Engineering Excellence 2026',
			role: 'Validation-tax guardrail',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.harness.io/state-of-engineering-excellence',
			note: 'Supports caution that generated-code volume can increase review, debugging, and coordination work; it does not mechanically reduce the score.'
		},
		{
			name: 'Lightrun State of AI-Powered Engineering 2026',
			role: 'Runtime-reliability guardrail',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://lightrun.com/ebooks/state-of-ai-powered-engineering-2026/',
			note: 'Used as reliability context for AI-generated code that still requires manual debugging after QA and staging.'
		},
		{
			name: 'CircleCI State of Software Delivery 2026',
			role: 'Delivery-system guardrail',
			cadence: 'Annual',
			type: 'telemetry report',
			url: 'https://circleci.com/blog/five-takeaways-2026-software-delivery-report/',
			note: 'Telemetry context for the gap between code volume and shipped value; not a positive delegation input.'
		},
		{
			name: 'Faros AI Engineering Report 2026',
			role: 'Telemetry-quality guardrail',
			cadence: 'Annual',
			type: 'telemetry report',
			url: 'https://www.faros.ai/blog/ai-acceleration-whiplash-takeaways',
			note: 'Customer telemetry supports reporting confidence and quality risk around AI acceleration rather than increasing the central score.'
		},
		{
			name: 'Debt Behind the AI Boom',
			role: 'Maintenance-risk guardrail',
			cadence: 'Periodic',
			type: 'research',
			url: 'https://arxiv.org/abs/2603.28592',
			note: 'Used as technical-debt context for AI-authored commits; it is not a delegation-share measure.'
		}
	],
	support: [
		{
			name: 'Salesforce State of Service 7th Edition',
			role: 'Scored cases-handled input',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.salesforce.com/news/stories/state-of-service-report-announcement-2025/',
			note: 'Provides the 30% current cases-handled-by-AI estimate; useful for delegation, but still survey-estimated rather than operational case telemetry.'
		},
		{
			name: 'Intercom Customer Service Transformation Report',
			role: 'Scored mature-deployment input',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.intercom.com/customer-transformation-report',
			note: 'Provides the 10% mature deployment signal where AI is integrated into support operations and working at scale.'
		},
		{
			name: 'Intercom Deflection Source',
			role: 'Held deflection input',
			cadence: 'Annual',
			type: 'operational estimate',
			note: 'Retained from the prior locked source because no clearer 2026 operational deflection metric with a stable resolved/no-escalation definition was found.'
		},
		{
			name: 'Sinch AI Production Paradox',
			role: 'Scored production-readiness input',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://sinch.com/news/sinch-releases-ai-production-paradox/',
			note: 'Provides the 62% live-production AI customer-communications signal, with rollback evidence retained as a reliability caveat.'
		},
		{
			name: 'Zendesk CX Trends 2026',
			role: 'Customer-facing AI context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://cxtrends.zendesk.com/',
			note: 'Useful for transparency and customer-expectation pressure, but not scored without a direct automated-resolution metric.'
		},
		{
			name: 'Gartner 2026 Customer Service AI Survey',
			role: 'Governance and operating-model context',
			cadence: 'Periodic',
			type: 'survey',
			url: 'https://www.gartner.com/en/newsroom/press-releases/2026-02-18-gartner-survey-finds-ninety-one-percent-of-customer-service-leaders-under-pressure-to-implement-ai-in-2026',
			note: 'Shows implementation pressure and labor-redesign context, not resolved-case delegation.'
		},
		{
			name: 'Hiver State of AI Customer Support 2026',
			role: 'Outcome-skepticism guardrail',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://hiverhq.com/reports/state-of-ai-customer-support-2026',
			note: 'Used as a counterweight to vendor-success narratives because it highlights limited resolution-time and cost-per-ticket improvements.'
		},
		{
			name: 'Stanford HAI AI Index Economy Chapter',
			role: 'Productivity context',
			cadence: 'Annual',
			type: 'synthesis report',
			url: 'https://hai.stanford.edu/assets/files/ai_index_report_2026_chapter_4_economy.pdf',
			note: 'Supports the domain importance through agent-assist productivity evidence, but is not an adoption or deflection-rate input.'
		}
	],
	credit: [
		{
			name: 'Upstart 2025 Form 10-K',
			role: 'Scored automation numerator',
			cadence: 'Annual',
			type: 'filing',
			url: 'https://www.sec.gov/Archives/edgar/data/1647639/000164763926000027/upst-20251231.htm',
			note: 'Provides the 91% fully automated platform-loan numerator used inside the personal-loan proxy; not treated as broad credit-market share.'
		},
		{
			name: 'TransUnion Q4 2025 CIIR',
			role: 'Scored market denominator',
			cadence: 'Quarterly',
			type: 'report',
			url: 'https://newsroom.transunion.com/q4-2025-ciir/',
			note: 'Provides the 42% fintech unsecured personal-loan origination denominator used with the Upstart platform automation rate.'
		},
		{
			name: 'McKinsey Credit Risk Survey',
			role: 'Conservative bank-decisioning hold',
			cadence: 'Annual',
			type: 'survey',
			note: 'Retained as a cautious bank credit-decisioning signal while direct 2026 bank underwriting deployment data remains unavailable.'
		},
		{
			name: 'Bank of England/FCA AI in UK Financial Services 2024',
			role: 'Automation and governance context',
			cadence: 'Periodic',
			type: 'survey',
			url: 'https://www.bankofengland.co.uk/report/2024/artificial-intelligence-in-uk-financial-services-2024',
			note: 'Broad financial-services AI survey used as a guardrail; useful for automated-decision context but not credit-specific enough to raise the bank score.'
		},
		{
			name: 'Bank of England 2026 AI Roundtables',
			role: 'Deployment-friction context',
			cadence: 'Periodic',
			type: 'roundtable',
			url: 'https://www.bankofengland.co.uk/minutes/2026/february/summary-of-ai-roundtables-feb-2026',
			note: 'Supports conservative treatment of regulated-bank deployment because risk, model validation, and governance functions remain cautious.'
		}
	],
	'medical-dx': [
		{
			name: 'FDA AI-Enabled Medical Device List',
			role: 'Scored device-count input',
			cadence: 'Continuous',
			type: 'database',
			url: 'https://www.fda.gov/medical-devices/software-medical-device-samd/artificial-intelligence-enabled-medical-devices',
			note: 'Provides the device-count signal, but device authorizations are not the same as clinical usage or diagnostic delegation.'
		},
		{
			name: 'Stanford HAI AI Index Medicine Chapter',
			role: 'Regulatory and evidence-quality context',
			cadence: 'Annual',
			type: 'synthesis report',
			url: 'https://hai.stanford.edu/ai-index/2026-ai-index-report/medicine',
			note: 'High-quality synthesis used to document regulatory growth and caveats; direct FDA data remains the operational count.'
		},
		{
			name: 'AMA 2026 Physician AI Sentiment Survey',
			role: 'Scored diagnosis-specific input',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.ama-assn.org/sites/ama-assn.org/files/2026-03/physician-ai-sentiment-report_0.pdf',
			note: 'Provides the 17% assistive-diagnosis value; broader physician AI use was excluded because it includes documentation and administrative workflows.'
		},
		{
			name: 'KLAS Global Imaging AI 2025',
			role: 'Scored imaging-adoption input',
			cadence: 'Annual',
			type: 'report',
			url: 'https://klasresearch.com/report/global-imaging-ai-2025-looking-at-adoption-and-usage-across-regions/3804',
			note: 'Provides the imaging AI adoption context, with caveats that it is organization-level, non-US, and not a physician-level diagnostic usage rate.'
		},
		{
			name: 'Doximity State of AI in Medicine 2026',
			role: 'Broad workflow context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.doximity.com/reports/state-of-ai-medicine-report/2026',
			note: 'Shows broad medical AI adoption, but is not diagnosis-specific enough to score as delegated diagnostic work.'
		},
		{
			name: 'OECD Scaling AI in Health',
			role: 'Health-system scale context',
			cadence: 'Periodic',
			type: 'policy report',
			url: 'https://www.oecd.org/content/dam/oecd/en/publications/reports/2026/04/scaling-artificial-intelligence-in-health_77610b12/a436e12d-en.pdf',
			note: 'Used to interpret the gap between local AI implementations and national-scale health-system deployment.'
		},
		{
			name: 'NVIDIA State of AI in Healthcare and Life Sciences 2026',
			role: 'Vendor-survey context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.nvidia.com/content/dam/en-zz/Solutions/lp/survey-report/healthcare-state-of-ai-report-2026-4559650-web.pdf',
			note: 'Broad healthcare and life-sciences adoption context, not a direct clinical diagnostic denominator.'
		},
		{
			name: 'Pathology AI Adoption Literature',
			role: 'Low-confidence scored hold',
			cadence: 'Periodic',
			type: 'research',
			url: 'https://journals.plos.org/digitalhealth/article?id=10.1371/journal.pdig.0001052',
			note: 'Supports a conservative pathology value because current evidence shows interest and localized use, not representative clinical deployment.'
		}
	],
	'legal-ai': [
		{
			name: 'Thomson Reuters 2026 AI in Professional Services',
			role: 'Scored organization-adoption input',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.thomsonreuters.com/content/dam/ewp-m/documents/thomsonreuters/en/pdf/reports/2026-ai-in-professional-services-report.pdf',
			note: 'Provides the 40% organization GenAI adoption input; use-case shares among GenAI users are retained as context rather than whole-market delegation rates.'
		},
		{
			name: 'Clio Solo and Small Firm Report 2026',
			role: 'Scored small-firm input',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.clio.com/about/press/2026-solo-small-firm-report/',
			note: 'Provides the solo and small-firm legal-work AI adoption signal, not an autonomous legal-output share.'
		},
		{
			name: 'Clio Mid-Sized Law Firms 2026',
			role: 'Firm-size adoption context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.clio.com/about/press/ai-is-reshaping-how-mid-sized-law-firms-scale-clio-reports/',
			note: 'Useful for a firm-size adoption curve, but not scored directly because it is mid-market rather than the current BigLaw/small-firm framing.'
		},
		{
			name: 'ABA Litigation and TAR TechReport 2024',
			role: 'Scored document-review anchor',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.americanbar.org/groups/law_practice/resources/tech-report/2024/2024-litigation-and-tar-techreport/',
			note: 'Older than the 2026 adoption sources, but still the cleanest workflow-specific denominator for AI-assisted search and TAR in litigation.'
		},
		{
			name: 'ABA Artificial Intelligence TechReport 2024',
			role: 'Governance and baseline context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.americanbar.org/groups/law_practice/resources/tech-report/2024/2024-artificial-intelligence-techreport/',
			note: 'Used as baseline and caution context; not fresh enough to anchor the 2026 legal score by itself.'
		}
	],
	hire: [
		{
			name: 'SHRM State of AI in HR 2026',
			role: 'Broad HR adoption guardrail',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.shrm.org/topics-tools/research/state-of-ai-hr-2026/full-report',
			note: 'Useful organization-level context for AI in HR, but the broad HR denominator is not treated as a direct hiring-delegation rate.'
		},
		{
			name: 'ICIMS/Aptitude AI Adoption in Talent Acquisition',
			role: 'Scored talent-acquisition workflow source',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.icims.com/company/newsroom/aiadoptionreport2026/',
			note: 'Provides the main 2026 hiring workflow values: AI use in talent acquisition, screening use-case adoption, and broad process coverage.'
		},
		{
			name: 'HireVue 2026 Global AI in Hiring Report',
			role: 'Candidate and vendor-context source',
			cadence: 'Annual',
			type: 'report',
			url: 'https://www.hirevue.com/resources/report/2026-global-ai-in-hiring-report',
			note: 'Retained as context for hiring-market AI pressure and candidate behavior; not used as a direct platform-reach denominator.'
		},
		{
			name: 'HireVue Q1 2024 Platform Volume Update',
			role: 'Low-confidence platform-reach proxy',
			cadence: 'Periodic',
			type: 'company update',
			url: 'https://www.hirevue.com/blog/hiring/quarterly-updates-from-hirevue-ceo',
			note: 'Supports the 80M/year reach proxy, but vendor assessment volume is not the same as overall application share.'
		}
	],
	education: [
		{
			name: 'Stanford HAI AI Index Education Chapter',
			role: 'Synthesis and source-discovery anchor',
			cadence: 'Annual',
			type: 'synthesis report',
			url: 'https://hai.stanford.edu/ai-index/2026-ai-index-report/education',
			note: 'Used as a high-quality education AI synthesis source; it supports schoolwork-use framing better than a narrow tutoring-only indicator.'
		},
		{
			name: 'Pew Research Center Teen AI Survey',
			role: 'Scored student-use input',
			cadence: 'Periodic',
			type: 'survey',
			url: 'https://www.pewresearch.org/internet/2026/02/24/how-teens-use-and-view-ai/',
			note: 'Provides the 54% U.S. teen schoolwork-help value, with lower all/most-use rates retained as high-delegation context.'
		},
		{
			name: 'OECD Digital Education Outlook 2026',
			role: 'Scored teacher-work input',
			cadence: 'Triennial',
			type: 'report',
			url: 'https://www.oecd.org/en/publications/2026/01/oecd-digital-education-outlook-2026_940e0dd8.html',
			note: 'Provides the 37% lower-secondary teacher AI-for-work value; it is not a direct grading-delegation measure.'
		},
		{
			name: 'EDUCAUSE AI and Higher-Education Work Report',
			role: 'Higher-ed work context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://www.educause.edu/research/2026/the-impact-of-ai-on-work-in-higher-education',
			note: 'Shows pervasive higher-ed AI work adoption, but is not scored until faculty or teaching-specific breakouts are locked.'
		},
		{
			name: 'EDUCAUSE Students and Technology Report 2026',
			role: 'Candidate student-source context',
			cadence: 'Annual',
			type: 'survey',
			url: 'https://library.educause.edu/resources/2026/3/2026-educause-students-and-technology-report-steady-through-change',
			note: 'Retained as a candidate source, but no numeric value is locked without full report extraction and citation.'
		},
		{
			name: 'Turnitin AI Writing / Clarity 2026',
			role: 'Scored student-output signal',
			cadence: 'Annual',
			type: 'report',
			url: 'https://www.prnewswire.com/news-releases/turnitin-data-shows-transparency-about-ai-use-benefits-students-and-educators-302695254.html',
			note: 'Provides the greater-than-80%-AI-written submissions signal, with detector and measurement caveats retained.'
		},
		{
			name: 'Gradescope / Turnitin Assessment Automation',
			role: 'Held grading-automation input',
			cadence: 'Annual',
			type: 'report',
			note: 'The AI-graded-assessments value is held from prior locked evidence because no stronger 2026 direct assessment-automation metric was found.'
		}
	]
};
