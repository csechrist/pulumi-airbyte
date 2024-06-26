// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceTwilio DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceTwilio = airbyte.getSourceTwilio({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceTwilio(args: GetSourceTwilioArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceTwilioResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceTwilio:getSourceTwilio", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceTwilio.
 */
export interface GetSourceTwilioArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceTwilio.
 */
export interface GetSourceTwilioResult {
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
 * SourceTwilio DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceTwilio = airbyte.getSourceTwilio({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceTwilioOutput(args: GetSourceTwilioOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceTwilioResult> {
    return pulumi.output(args).apply((a: any) => getSourceTwilio(a, opts))
}

/**
 * A collection of arguments for invoking getSourceTwilio.
 */
export interface GetSourceTwilioOutputArgs {
    sourceId: pulumi.Input<string>;
}
