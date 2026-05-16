# Education 2026 Evidence Extraction

Status: implemented for the 2026 Q2 source refresh.
Prepared: 2026-05-16.

## Current Scoring Contract

Current `education` score: 35.8.

Configured indicators:

- `Students Using AI for Schoolwork`: 40% weight.
- `AI-Graded Assessments`: 20% weight.
- `Teachers Using AI for Work`: 30% weight.
- `Student Papers 80%+ AI-Written`: 10% weight.

Current observations:

- `Students Using AI for Schoolwork`: 54.0%, freshness 2026.
- `AI-Graded Assessments`: 8.0%, freshness 2025.
- `Teachers Using AI for Work`: 37.0%, freshness 2024.
- `Student Papers 80%+ AI-Written`: 15.0%, freshness 2026.

The current score calculates as:

```text
0.40 * 54.0 + 0.20 * 8.0 + 0.30 * 37.0 + 0.10 * 15.0 = 35.8
```

The implemented 2026 Q2 contract broadens the student indicator from tutoring to schoolwork, keeps grading separate, and adds a low-weight student-output signal with detector caveats.

## Extracted Candidate Sources

### Stanford HAI 2026 AI Index Education Chapter

Source: https://hai.stanford.edu/ai-index/2026-ai-index-report/education
Published: 2026.
Measurement period: 2025/2026 synthesis.
Evidence grade: B for broad workflow penetration.
Confidence: high.

Relevant values:

- Four out of five U.S. high school and college students use AI for schoolwork.
- Only half of middle and high schools have AI policies.
- Only 6% of teachers say school AI policies are clear.
- Common student use cases include research, essay editing, and brainstorming.

Recommendation: use as a high-quality synthesis source and source-discovery anchor. It is a better match for `Students Using AI for Schoolwork` than for `Students Using AI Tutors`. Do not score it as "tutoring" unless the domain intentionally narrows to tutoring.

### Pew Research Center: How Teens Use and View AI

Source: https://www.pewresearch.org/internet/2026/02/24/how-teens-use-and-view-ai/
Published: 2026-02-24.
Measurement period: survey conducted 2025-09-25 to 2025-10-09.
Evidence grade: B for student workflow penetration; C for depth of delegated schoolwork.
Confidence: high.

Relevant values:

- 64% of teens report using chatbots.
- 54% of U.S. teens say they have used chatbots for schoolwork help.
- 10% of teens say they do all or most schoolwork with chatbot help.
- 21% say chatbots help with some schoolwork.
- 23% say chatbots help with a little schoolwork.
- 48% have used chatbots to research a topic for school.
- 43% have used them to solve a math problem.
- 35% have used them to edit something they wrote.
- 25% say chatbots are extremely or very helpful for schoolwork, and another 25% say somewhat helpful.

Recommendation: use Pew as the strongest K-12 depth source. If the scored indicator becomes `Students Using AI for Schoolwork`, use 54% as broad use and 31% as "some or more" depth. The 10% all/most value is a useful high-delegation sub-signal.

### OECD Digital Education Outlook 2026

Source: https://www.oecd.org/en/publications/2026/01/oecd-digital-education-outlook-2026_940e0dd8.html
Published: 2026-01-19.
Measurement period: TALIS 2024 and OECD synthesis.
Evidence grade: B for teacher workflow penetration.
Confidence: high.

Relevant values:

- 37% of lower-secondary teachers used AI for their job in 2024.
- 57% agree AI helps write or improve lesson plans.
- 72% believe AI can harm academic integrity by allowing students to pass off work as their own.

Recommendation: use to refresh the teacher/faculty workflow indicator if we broaden it to `Teachers Using AI for Work` or `Teachers Using AI in Instructional Work`. It is not a direct grading-delegation measure.

### EDUCAUSE: The Impact of AI on Work in Higher Education

Source: https://www.educause.edu/research/2026/the-impact-of-ai-on-work-in-higher-education
Published: 2026-01-12.
Measurement period: survey conducted 2025-09-29 to 2025-10-13.
Evidence grade: B for higher-ed workforce penetration.
Confidence: medium.

Relevant values:

- 1,960 responses met inclusion criteria.
- 94% reported using AI tools for work within the past six months.
- Among recent users, 73% used AI tools daily or weekly for work-related tasks.
- 54% used AI tools for eight or more types of work-related tasks in the past six months.
- 56% used AI tools not provided by their institutions.
- Respondents selected `Creating assessments` as a promising opportunity at 29%.
- Respondents selected `Evaluating student work` as a promising opportunity at 16%.

Recommendation: use as context for higher-ed faculty/staff work adoption and governance, not as a direct `Faculty Using AI in Teaching` score unless role-specific faculty/teaching breakouts are extracted. It supports the claim that AI is pervasive in higher-ed work, but it is not a clean teaching-only measure.

### EDUCAUSE Students and Technology Report 2026

Source: https://library.educause.edu/resources/2026/3/2026-educause-students-and-technology-report-steady-through-change
Published: 2026-03-23.
Measurement period: 2026 report.
Evidence grade: pending.
Confidence: medium.

Relevant values:

- Public landing page confirms the report covers student perspectives on generative AI, workforce preparation, course experiences, and student support.
- Full report access is member-gated in the current browsing context.

Recommendation: keep in the ledger as a candidate, but do not lock a numeric value until the full report can be accessed and cited.

### Turnitin AI Writing / Clarity 2026

Source: https://www.prnewswire.com/news-releases/turnitin-data-shows-transparency-about-ai-use-benefits-students-and-educators-302695254.html
Published: 2026-02-24.
Measurement period: since October 2025 for latest AI detection tool data; first three months of Clarity use for prompting analysis.
Evidence grade: C for student output share; B for platform process data.
Confidence: medium.

Relevant values:

- Since October 2025, approximately 15% of essay submissions had greater than 80% AI-generated writing.
- This compares with an average of 3% when Turnitin launched its original AI detector in April 2023.
- In the first three months of Turnitin Clarity use, 29% of student prompts asked for review, judgment, or other feedback.
- 94% of students wrote their own prompts rather than using pre-written prompt suggestions.
- 36% of prompts in the feedback category were considered effective.

Recommendation: retain `Student Papers 80%+ AI-Written` as a displayed student-output signal, with explicit detector caveats. Do not mix it into `AI-Graded Assessments`; it measures student production, not educator assessment delegation.

## Implemented Education Source Lock

Current v2 scoring inputs and retained candidates:

| Indicator | Role | Evidence grade | Confidence | Notes |
| --- | --- | --- | --- | --- |
| Students Using AI for Schoolwork | score input | B | high | Pew 54% teen schoolwork help is used as the current K-12 depth anchor; HAI broad 80% remains context |
| AI-Graded Assessments | score input if direct source found | A/B | low | No refreshed direct assessment-automation source locked yet |
| Faculty/Teachers Using AI in Instructional Work | score input | B | medium | OECD 37% lower-secondary teacher work use; EDUCAUSE 94% higher-ed work use is broader than teaching |
| Teachers Using AI for Grading | context or pending | B | low | Need direct actual-use measure; EDUCAUSE 16% is opportunity, not usage |
| Student Papers 80%+ AI-Written | low-weight output-share signal | C | medium | Turnitin 15%; detector caveats and submission-population bias |

Implementation decision:

- Rename `Students Using AI Tutors` to `Students Using AI for Schoolwork`.
- Keep `AI-Graded Assessments` separate and do not backfill it with student AI-writing or detection data.
- Use Pew as the current K-12 schoolwork score input and HAI as broad cross-level context.
- Use OECD for teacher work use unless a better faculty-teaching-specific source is found.
- Use Turnitin as a low-weight student-output share signal with caveats, not as evidence of grading delegation.
