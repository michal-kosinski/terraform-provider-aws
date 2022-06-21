---
subcategory: "Kendra"
layout: "aws"
page_title: "AWS: aws_kendra_experience"
description: |-
  Terraform resource for managing an AWS Kendra Experience.
---

# Resource: aws_kendra_experience

Terraform resource for managing an AWS Kendra Experience.

## Example Usage

### Basic Usage

```terraform
resource "aws_kendra_experience" "example" {
  index_id    = aws_kendra_index.example.id
  description = "My Kendra Experience"
  name        = "example"
  role_arn    = aws_iam_role.example.arn

  configuration {
    content_source_configuration {
      direct_put_content = true
    }
    user_identity_configuration {
      identity_attribute_name = "12345ec453-1546651e-79c4-4554-91fa-00b43ccfa245"
    }
  }
}
```

## Argument Reference

The following arguments are required:

* `index_id` - (Required, Forces new resource) The identifier of the index for your Amazon Kendra experience.
* `name` - (Required) A name for your Amazon Kendra experience.
* `role_arn` - (Required) The Amazon Resource Name (ARN) of a role with permission to access `Query API`, `QuerySuggestions API`, `SubmitFeedback API`, and `AWS SSO` that stores your user and group information. For more information, see [IAM roles for Amazon Kendra](https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).

The following arguments are optional:

* `description` - (Optional, Forces new resource if removed) A description for your Amazon Kendra experience.
* `configuration` - (Optional) Configuration information for your Amazon Kendra experience. [Detailed below](#configuration).

### `configuration`

~> **NOTE:** At least one of `content_source_configuration` or `user_identity_configuration` must be configured.

The `configuration` configuration block supports the following arguments:

* `content_source_configuration` - (Optional) The identifiers of your data sources and FAQs. Or, you can specify that you want to use documents indexed via the `BatchPutDocument API`. [Detailed below](#content_source_configuration).
* `user_identity_configuration` - (Optional) The AWS SSO field name that contains the identifiers of your users, such as their emails. [Detailed below](#user_identity_configuration).

### `content_source_configuration`

The `content_source_configuration` configuration block supports the following arguments:

* `data_source_ids` - (Optional) The identifiers of the data sources you want to use for your Amazon Kendra experience. Maximum number of 100 items.
* `direct_put_content` - (Optional) Whether to use documents you indexed directly using the `BatchPutDocument API`. Defaults to `false`.
* `faq_ids` - (Optional) The identifier of the FAQs that you want to use for your Amazon Kendra experience. Maximum number of 100 items.

### `user_identity_configuration`

The `user_identity_configuration` configuration block supports the following argument:

* `identity_attribute_name` - (Required) The AWS SSO field name that contains the identifiers of your users, such as their emails.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifiers of the experience and index separated by a slash (`/`).
* `arn` - ARN of the Experience.
* `endpoints` - Shows the endpoint URLs for your Amazon Kendra experiences. The URLs are unique and fully hosted by AWS.
    * `endpoint` - The endpoint of your Amazon Kendra experience.
    * `endpoint_type` - The type of endpoint for your Amazon Kendra experience.
* `experience_id` - The unique identifier of the experience.
* `status` - The current processing status of your Amazon Kendra experience.

## Timeouts

`aws_kendra_experience` provides the following [Timeouts](https://www.terraform.io/docs/configuration/blocks/resources/syntax.html#operation-timeouts) configuration options:

* `create` - (Optional, Default: `30m`)
* `update` - (Optional, Default: `30m`)
* `delete` - (Optional, Default: `30m`)

## Import

Kendra Experience can be imported using the unique identifiers of the experience and index separated by a slash (`/`) e.g.,

```
$ terraform import aws_kendra_experience.example 1045d08d-66ef-4882-b3ed-dfb7df183e90/b34dfdf7-1f2b-4704-9581-79e00296845f
```
