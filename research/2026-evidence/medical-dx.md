# Medical Diagnosis 2026 Evidence Extraction

Status: source refresh notes only; no score update yet.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `medical-dx` score: 25.1.

Configured indicators:

- `FDA-Cleared Diagnostic AI Devices`: 10% weight.
- `Radiology AI Adoption`: 35% weight.
- `AI-Assisted Diagnosis Rate`: 30% weight.
- `Pathology AI Adoption`: 25% weight.

Current observations:

- `FDA-Cleared Diagnostic AI Devices`: 1,016 devices, freshness 2025.
- `Radiology AI Adoption`: 30%, freshness 2025.
- `AI-Assisted Diagnosis Rate`: 12%, freshness 2025.
- `Pathology AI Adoption`: 10%, freshness 2025.
- Display-only context includes autonomous diabetic retinopathy screening, European radiologist clinical use, health systems with imaging AI deployed, and physicians using AI for diagnosis.

The current contract is useful but mixes three different quantities: regulatory capacity, professional workflow adoption, and diagnostic delegation. The 2026 refresh should keep those quantities separate and avoid treating benchmark diagnostic performance as clinical deployment.

## Extracted Candidate Sources

### FDA AI-Enabled Medical Device List

Source: https://www.fda.gov/medical-devices/software-medical-device-samd/artificial-intelligence-enabled-medical-devices
Published: continuously updated FDA list.
Measurement period: accessed 2026-05-16.
Evidence grade: A for authorization inventory; C for clinical usage.
Confidence: high for count; low for usage interpretation.

Local count from FDA CSV download on 2026-05-16:

- 1,430 rows in the current AI-enabled medical device list.
- 331 devices had a final decision date in 2025.
- 1,094 listed devices have `Radiology` as the lead panel.
- 9 listed devices have `Pathology` as the lead panel.
- 9 listed devices have `Clinical Chemistry` as the lead panel.
- 10 listed devices have `Ophthalmic` as the lead panel.

Important caveats:

- The FDA says the list is not comprehensive; it is based primarily on AI-related terms in public authorization summaries and classifications.
- The list identifies devices authorized for marketing, not installed devices, clinician use, patient encounters, or autonomous diagnosis.
- The current score's normalization max of 1,200 will saturate if the full 1,430-device count is used unchanged.

Recommendation: refresh the device count from the FDA CSV, but rename the indicator to `FDA AI-Enabled Medical Devices` or use a diagnostic/radiology-relevant subset. Do not interpret count growth as clinical delegation growth without a usage denominator.

### Stanford HAI 2026 AI Index Medicine Chapter

Source: https://hai.stanford.edu/ai-index/2026-ai-index-report/medicine
Published: 2026.
Measurement period: 2025 synthesis.
Evidence grade: B for field context.
Confidence: high.

Relevant values:

- HAI reports 258 FDA AI medical devices authorized in 2025.
- HAI reports that most 2025 authorizations used pathways that do not require new clinical trials.
- HAI reports only 2.4% of devices with clinical studies were supported by randomized-trial data.
- Clinical-note automation saw broad 2025 adoption; physicians at multiple hospital systems reported up to 83% less time writing notes, and one hospital system reported 112% ROI.
- A multi-agent diagnostic system scored 85.5% on complex published case studies versus 20% for unaided physicians.
- AI-generated summaries appeared at the top of 84%-92% of health-related Google searches.

Recommendation: use HAI as a high-quality context and source-discovery anchor. The 258-device figure may differ from a live FDA CSV count because of report timing or counting definitions. Use direct FDA as the operational count and HAI to document regulatory and evidence-quality caveats.

### AMA 2026 Physician Survey on Augmented Intelligence

Source: https://www.ama-assn.org/sites/ama-assn.org/files/2026-03/physician-ai-sentiment-report_0.pdf
Published: 2026-03.
Measurement period: survey conducted 2026-01-15 to 2026-02-02.
Evidence grade: A for U.S. physician use-case adoption.
Confidence: high.

Relevant values:

- Survey of 1,692 physicians across specialties, practice settings, and career stages.
- 81% of physicians reported awareness or use of AI in their practice in 2026.
- 72% reported incorporating one or more AI use cases into practice.
- The average number of AI use cases per physician increased from 1.1 in 2023 to 2.3 in 2026.
- 39% use AI for summaries of medical research and standards of care.
- 30% use AI for creation of discharge instructions, care plans, or progress notes.
- 28% use AI for documentation of billing codes, medical charts, or visit notes.
- 28% use AI for generation of chart summaries.
- 17% use AI for assistive diagnosis, up from 12% in 2024.
- Among physicians who considered the use case relevant, 45% already use or expect to use assistive diagnosis by the end of 2026.
- 70% expect AI to offload or replace some clinical tasks, while 73% expect administrative workload reduction.
- Nearly half would never or rarely want patients using AI to interpret pathology results or radiology reports/images without physician involvement.

Recommendation: use the 17% assistive-diagnosis value as the leading candidate to refresh `AI-Assisted Diagnosis Rate`. Keep the broader 81%/72% physician-use values out of the diagnosis score because they include administrative and documentation workflows.

### Doximity 2026 State of AI in Medicine

Source: https://www.doximity.com/reports/state-of-ai-medicine-report/2026
Published: 2026.
Measurement period: surveys in 2025-03 to 2025-04 and 2025-11 to 2026-01.
Evidence grade: B for broad physician AI adoption.
Confidence: medium.

Relevant values:

- Survey responses from 3,151 U.S. physicians across 15 specialties.
- 94% of physicians surveyed are currently using AI or are interested in doing so.
- 54% reported currently using AI in clinical practice overall.
- Adoption rose from 47% in March-April 2025 to 63% in the November 2025-January 2026 wave.
- 37% of all physicians surveyed reported using AI at least daily.
- Among AI users, 69% reported daily use.
- The most common use case was literature search at 35% of physicians in January 2026, followed by ambient documentation at 29%.

Recommendation: use as a broad medical workflow adoption source, not as a diagnostic delegation input. Doximity strengthens the case that physician AI adoption is broadening, but the available public summary does not give a diagnosis-specific use rate.

### KLAS Global Imaging AI 2025

Source: https://klasresearch.com/report/global-imaging-ai-2025-looking-at-adoption-and-usage-across-regions/3804
Published: 2025.
Measurement period: 2025 report.
Evidence grade: B for imaging AI adoption outside the U.S.
Confidence: medium.

Relevant values:

- KLAS interviewed 369 organizations across Asia/Oceania, Canada, Europe, Latin America, and the Middle East/Africa.
- Nearly 50% of the organizations interviewed, excluding the United States, now use at least one imaging AI solution.
- Organizations reported using 108 different AI vendors.
- The report covers pixel AI, operational/reporting AI, and AI platforms.
- KLAS notes that customer satisfaction is not explored in this report.

Recommendation: use as refreshed radiology/imaging adoption context if the methodology accepts non-U.S. organization-level survey data. It is not a clean physician-level usage rate and excludes the U.S.

### OECD Scaling AI in Health

Source: https://www.oecd.org/content/dam/oecd/en/publications/reports/2026/04/scaling-artificial-intelligence-in-health_77610b12/a436e12d-en.pdf
Published: 2026-04.
Measurement period: OECD expert-group synthesis.
Evidence grade: B for policy and scale context.
Confidence: medium-high.

Relevant values:

- OECD reports AI is used in health systems across all OECD countries interviewed.
- National-level scale-up remains limited, with only 10% for medical imaging applications.
- OECD highlights fragmented data foundations, non-aligned policies, and governance barriers as constraints on scaling.
- The report frames medical imaging and administrative automation as the dominant areas to date.

Recommendation: use OECD to interpret the difference between local implementation and national-scale adoption. It should be a context source, not a direct score input unless the next methodology adds a health-system-scale indicator.

### NVIDIA State of AI in Healthcare and Life Sciences 2026

Source: https://www.nvidia.com/content/dam/en-zz/Solutions/lp/survey-report/healthcare-state-of-ai-report-2026-4559650-web.pdf
Published: 2026.
Measurement period: 2025/2026 survey.
Evidence grade: C for medical diagnosis scoring; B for industry adoption context.
Confidence: medium-low because it is a vendor survey.

Relevant values:

- 70% of respondents said their organizations actively use AI, up from 63% in the prior survey.
- 42% said they use AI to support clinical decision making.
- The payers/providers segment reported 39% clinical decision support and 40% NLP in clinical documentation.
- The medtech/tools/diagnostics segment reported 61% medical imaging and 34% diagnostic testing including disease diagnosis and risk prediction.
- 47% said their organizations are actively using or assessing AI agents.

Recommendation: use as broad industry context only. The respondent base is healthcare and life sciences organizations, not necessarily clinical diagnostic users, and the vendor source needs a caveat.

### Pathology AI Adoption Sources

Sources:

- https://www.sciencedirect.com/science/article/pii/S2153353925001129
- https://journals.plos.org/digitalhealth/article?id=10.1371/journal.pdig.0001052

Published: 2025.
Measurement period: 2024/2025 surveys.
Evidence grade: C for global/clinical pathology adoption.
Confidence: low to medium.

Relevant values:

- The global pathology survey had 268 respondents from 23 countries, 65% from the United States. Actual use was limited: 31% reported rare use and 29% no use at all.
- In that global survey, AI was used mainly for document drafting, research, and administration; diagnostic use was minimal. Only 10% reported clear institutional AI guidelines.
- The PLOS Digital Health AI-assisted diagnostic systems survey included 224 pathologists in China. 37.9% had used AIADS and over 80% supported AIADS in clinical diagnostics, but diagnostic accuracy remained the main concern.
- As of the FDA count above, pathology remains a tiny share of U.S. AI-enabled device authorizations.

Recommendation: keep `Pathology AI Adoption` as hold/pending unless a representative clinical pathology deployment survey is found. The latest sources show interest and localized use, but not a clean global or U.S. clinical adoption percentage comparable to radiology.

## Proposed Medical-Dx Source Lock

Proposed v2 scoring candidates:

| Indicator | Suggested role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| FDA AI-Enabled Medical Devices | low-weight capacity input | A/C | high for count | Direct FDA count is 1,430 as of 2026-05-16; count saturates current normalization and is not usage |
| Radiology or Imaging AI Adoption | score input | B | medium | KLAS nearly 50% non-U.S. organization adoption; need U.S. or physician-level comparator if possible |
| AI-Assisted Diagnosis Rate | score input | A | high | AMA 2026 assistive diagnosis is 17%; broader physician AI use should not be substituted |
| Pathology AI Adoption | hold or low confidence score input | C | low | Current sources show cautious and uneven adoption; representative direct deployment source still missing |
| Clinical Note AI Automation | context only | B | high | Strong workflow automation signal but outside diagnosis |
| AI Diagnostic Benchmark Performance | context only | B | high | Important capability signal but not clinical delegation |

Near-term decision:

- Use AMA 17% to refresh diagnosis-specific physician use if scoring proceeds before stronger sources are found.
- Refresh FDA count from direct FDA CSV, but decide whether the indicator should be all AI-enabled medical devices or a diagnostic/radiology subset.
- Keep KLAS as radiology/imaging adoption context and look for a U.S. or radiologist-level update before changing the radiology score.
- Keep pathology pending; do not replace it with digital pathology readiness or general pathology ChatGPT use.
- Keep clinical-note automation, Google health AI-overview prevalence, and diagnostic benchmark performance out of the diagnosis score unless the domain is broadened beyond clinical diagnosis.
