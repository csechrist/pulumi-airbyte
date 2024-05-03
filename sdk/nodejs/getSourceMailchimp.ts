// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceMailchimp DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceMailchimp = airbyte.getSourceMailchimp({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceMailchimp(args: GetSourceMailchimpArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceMailchimpResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceMailchimp:getSourceMailchimp", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceMailchimp.
 */
export interface GetSourceMailchimpArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceMailchimp.
 */
export interface GetSourceMailchimpResult {
    /**
     * The values required to configure the source. Parsed as JSON.
     */
    readonly configuration: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly sourceId: string;
    readonly sourceType: string;
    readonly workspaceId: string;
}
/**
 * SourceMailchimp DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceMailchimp = airbyte.getSourceMailchimp({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceMailchimpOutput(args: GetSourceMailchimpOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceMailchimpResult> {
    return pulumi.output(args).apply((a: any) => getSourceMailchimp(a, opts))
}

/**
 * A collection of arguments for invoking getSourceMailchimp.
 */
export interface GetSourceMailchimpOutputArgs {
    sourceId: pulumi.Input<string>;
}