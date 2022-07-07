# Validate JSON schema

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/steps-validate-json-schema?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/steps-validate-json-schema/releases)

Run JSON schema validation.

<details>
<summary>Description</summary>

<nil>
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://devcenter.bitrise.io/steps-and-workflows/steps-and-workflows-index/).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `schema_url` | URL of the JSON schema to use.  Use `file://` prefix for local, while `http` prefix for remote schemas. | required |  |
| `yaml_path` | Path of the .yml file to validate. | required |  |
| `warning_patterns` | Mark the schema validation issues, matching any of the provided regexp patterns, as warning.  Provide a newline (\n) separated list of patterns. |  |  |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/steps-validate-json-schema/pulls) and [issues](https://github.com/bitrise-steplib/steps-validate-json-schema/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://devcenter.bitrise.io/bitrise-cli/run-your-first-build/).

Learn more about developing steps:

- [Create your own step](https://devcenter.bitrise.io/contributors/create-your-own-step/)
- [Testing your Step](https://devcenter.bitrise.io/contributors/testing-and-versioning-your-steps/)
