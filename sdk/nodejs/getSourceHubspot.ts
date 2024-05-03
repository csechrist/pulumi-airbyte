// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceHubspot DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceHubspot = airbyte.getSourceHubspot({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceHubspot(args: GetSourceHubspotArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceHubspotResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceHubspot:getSourceHubspot", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceHubspot.
 */
export interface GetSourceHubspotArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceHubspot.
 */
export interface GetSourceHubspotResult {
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
 * SourceHubspot DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceHubspot = airbyte.getSourceHubspot({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceHubspotOutput(args: GetSourceHubspotOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceHubspotResult> {
    return pulumi.output(args).apply((a: any) => getSourceHubspot(a, opts))
}

/**
 * A collection of arguments for invoking getSourceHubspot.
 */
export interface GetSourceHubspotOutputArgs {
    sourceId: pulumi.Input<string>;
}