# AWS CRON

List all scheduled jobs (lambda, fargate) via Event Bridge.

## Example of command output

```bash
┌─────────────────────┬──────────┬──────────────────────────────────────────────────────────────┐
│      Schedule       │  Status  │                             Name                             │
├─────────────────────┼──────────┼──────────────────────────────────────────────────────────────┤
│ cron(0 7 * * ? *)   │ ENABLED  │ automation-test-seed-data                                    │
│                     │          │ Fires once a day to trigger the automation-test-seed-data    │
│                     │          │ job                                                          │
├─────────────────────┼──────────┼──────────────────────────────────────────────────────────────┤
│ cron(0 8 * * ? *)   │ ENABLED  │ automation-test                                              │
│                     │          │ Fires once a day to trigger the automation-test job          │
├─────────────────────┼──────────┼──────────────────────────────────────────────────────────────┤
│ rate(5 minutes)     │ ENABLED  │ messages-fanout                                              │
│                     │          │ Triggers job to deliver pending messages                     │
├─────────────────────┼──────────┼──────────────────────────────────────────────────────────────┤
│ cron(0 1 * * ? *)   │ DISABLED │ send-report                                                  │
│                     │          │ Fires once a day to trigger job that send daily report       │
└─────────────────────┴──────────┴──────────────────────────────────────────────────────────────┘
```

## Usage

```bash
aws-cron --help

# show list of cron jobs
aws-cron
aws-cron --output=table
aws-cron -o=text

# todo: show cli version
aws-cron version

# todo: trigger cronjob
aws-cron <name/arn?> trigger
```

## Release

```bash
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0

```