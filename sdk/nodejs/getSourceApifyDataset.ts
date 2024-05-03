// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceApifyDataset DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceApifydataset = airbyte.getSourceApifyDataset({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceApifyDataset(args: GetSourceApifyDatasetArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceApifyDatasetResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceApifyDataset:getSourceApifyDataset", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceApifyDataset.
 */
export interface GetSourceApifyDatasetArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceApifyDataset.
 */
export interface GetSourceApifyDatasetResult {
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
 * SourceApifyDataset DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceApifydataset = airbyte.getSourceApifyDataset({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceApifyDatasetOutput(args: GetSourceApifyDatasetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceApifyDatasetResult> {
    return pulumi.output(args).apply((a: any) => getSourceApifyDataset(a, opts))
}

/**
 * A collection of arguments for invoking getSourceApifyDataset.
 */
export interface GetSourceApifyDatasetOutputArgs {
    sourceId: pulumi.Input<string>;
}