# Medical Diagnosis: 2026 Q4 Evidence Notes

Run `2026-q4`, method `delegation-curve-v2`. Prior notes: [2026 Q2](../2026-q2/medical-dx.md).

Domain score 32.3 → **33.2**, from the FDA device count.

## Lock Decisions

| Indicator | Action | Value | Source | Window | Grade / confidence |
| --- | --- | ---: | --- | --- | --- |
| FDA AI-Enabled Medical Devices | refresh | 1,614 devices | [FDA AI-enabled device list](https://www.fda.gov/medical-devices/software-medical-device-samd/artificial-intelligence-enabled-medical-devices) (CSV, current as of 2026-09-04) | Final decisions through 2026-06-29 | A / high |
| Radiology or Imaging AI Adoption | hold | 50% | [KLAS Global Imaging AI 2025](https://klasresearch.com/report/global-imaging-ai-2025-looking-at-adoption-and-usage-across-regions/3804) | 2025 | B / medium |
| AI-Assisted Diagnosis Rate | hold | 17% | AMA 2026 physician survey | Jan 15-Feb 2, 2026 | A / high |
| Pathology AI Adoption | hold | 10% | Conservative hold | 2024-2025 | C / low |

## Notes

- **FDA.** The CSV has 1,614 rows, 184 more than in May; the HTML list agrees. The 0-2000 normalization cap now maps the count to 80.7 and will saturate around 2027 regardless of clinical use.
- **Radiology.** KLAS lists a [US & Global Imaging AI 2026](https://klasresearch.com/report/us-and-global-imaging-ai-2026/4013) report as upcoming for October 2026. If it is out at publication, decide whether a U.S.-inclusive figure is a refresh or a population change. The UK's [2025 radiology workforce census](https://www.rcr.ac.uk/media/n1fjvrv4/rcr-2025-clinical-radiology-workforce-census-report.pdf) (75% of departments using AI, full response) is a stronger survey but UK-only; it is recorded as a guardrail.
- **Pathology.** Newer CAP and KLAS figures measure digital pathology (for example 15% of practices reading slides digitally for primary diagnosis), not AI. They bound the plausible range but do not support a value.

## Candidates For Methodology v3

Medicare fee-for-service claims (CMS public data API, data year 2024) give reproducible utilization shares for AI-specific services:

- **FFR-CT:** 13.9% of coronary CT angiography studies had an AI-derived FFR-CT billed (29,270 of 210,289), up from 8.8% in 2023 and 6.8% in 2022.
- **Autonomous AI diabetic retinal exams (CPT 92229):** 25.4% of point-of-care or remote retinal exams but 0.15% of all retinal imaging; the denominator decides the value.
- **AI coronary plaque analysis:** about 1.1% of coronary CT angiography studies.
