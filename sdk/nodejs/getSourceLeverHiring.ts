// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceLeverHiring DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceLeverhiring = airbyte.getSourceLeverHiring({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceLeverHiring(args: GetSourceLeverHiringArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceLeverHiringResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceLeverHiring:getSourceLeverHiring", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceLeverHiring.
 */
export interface GetSourceLeverHiringArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceLeverHiring.
 */
export interface GetSourceLeverHiringResult {
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
 * SourceLeverHiring DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceLeverhiring = airbyte.getSourceLeverHiring({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceLeverHiringOutput(args: GetSourceLeverHiringOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceLeverHiringResult> {
    return pulumi.output(args).apply((a: any) => getSourceLeverHiring(a, opts))
}

/**
 * A collection of arguments for invoking getSourceLeverHiring.
 */
export interface GetSourceLeverHiringOutputArgs {
    sourceId: pulumi.Input<string>;
}
