// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceLinkedinAds DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceLinkedinads = airbyte.getSourceLinkedinAds({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceLinkedinAds(args: GetSourceLinkedinAdsArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceLinkedinAdsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceLinkedinAds:getSourceLinkedinAds", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceLinkedinAds.
 */
export interface GetSourceLinkedinAdsArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceLinkedinAds.
 */
export interface GetSourceLinkedinAdsResult {
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
 * SourceLinkedinAds DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceLinkedinads = airbyte.getSourceLinkedinAds({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceLinkedinAdsOutput(args: GetSourceLinkedinAdsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceLinkedinAdsResult> {
    return pulumi.output(args).apply((a: any) => getSourceLinkedinAds(a, opts))
}

/**
 * A collection of arguments for invoking getSourceLinkedinAds.
 */
export interface GetSourceLinkedinAdsOutputArgs {
    sourceId: pulumi.Input<string>;
}
