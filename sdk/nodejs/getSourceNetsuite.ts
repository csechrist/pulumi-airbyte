// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceNetsuite DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceNetsuite = airbyte.getSourceNetsuite({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceNetsuite(args: GetSourceNetsuiteArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceNetsuiteResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceNetsuite:getSourceNetsuite", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceNetsuite.
 */
export interface GetSourceNetsuiteArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceNetsuite.
 */
export interface GetSourceNetsuiteResult {
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
 * SourceNetsuite DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceNetsuite = airbyte.getSourceNetsuite({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceNetsuiteOutput(args: GetSourceNetsuiteOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceNetsuiteResult> {
    return pulumi.output(args).apply((a: any) => getSourceNetsuite(a, opts))
}

/**
 * A collection of arguments for invoking getSourceNetsuite.
 */
export interface GetSourceNetsuiteOutputArgs {
    sourceId: pulumi.Input<string>;
}
