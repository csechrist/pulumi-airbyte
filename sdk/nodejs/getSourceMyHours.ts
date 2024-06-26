// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * SourceMyHours DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceMyhours = airbyte.getSourceMyHours({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceMyHours(args: GetSourceMyHoursArgs, opts?: pulumi.InvokeOptions): Promise<GetSourceMyHoursResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("airbyte:index/getSourceMyHours:getSourceMyHours", {
        "sourceId": args.sourceId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSourceMyHours.
 */
export interface GetSourceMyHoursArgs {
    sourceId: string;
}

/**
 * A collection of values returned by getSourceMyHours.
 */
export interface GetSourceMyHoursResult {
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
 * SourceMyHours DataSource
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as airbyte from "@pulumi/airbyte";
 *
 * const mySourceMyhours = airbyte.getSourceMyHours({
 *     sourceId: "...my_source_id...",
 * });
 * ```
 */
export function getSourceMyHoursOutput(args: GetSourceMyHoursOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSourceMyHoursResult> {
    return pulumi.output(args).apply((a: any) => getSourceMyHours(a, opts))
}

/**
 * A collection of arguments for invoking getSourceMyHours.
 */
export interface GetSourceMyHoursOutputArgs {
    sourceId: pulumi.Input<string>;
}
