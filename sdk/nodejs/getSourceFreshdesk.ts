// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceFreshdesk DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceFreshdesk = airbyte.getSourceFreshdesk({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceFreshdesk(args: GetSourceFreshdeskArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceFreshdeskResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceFreshdesk:getSourceFreshdesk", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceFreshdesk.
 */
export interface GetSourceFreshdeskArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceFreshdesk.
 */
export interface GetSourceFreshdeskResult {
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
 * SourceFreshdesk DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceFreshdesk = airbyte.getSourceFreshdesk({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceFreshdeskOutput(args: GetSourceFreshdeskOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceFreshdeskResult> {
    return pulumi.output(args).apply((a: any) => getSourceFreshdesk(a, opts))
}

/**
 * A collection of arguments for invoking getSourceFreshdesk.
 */
export interface GetSourceFreshdeskOutputArgs {
    sourceId: pulumi.Input<string>;
}